package internal

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Firestore struct {
	client *firestore.Client
}

func NewFirestore() (*Firestore, error) {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	config := &firebase.Config{ProjectID: "mosho-166cd"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		return nil, err
	}

	log.Println("Firestore: Initialized")

	return &Firestore{client}, nil
}

func (f *Firestore) Next() (*Cmd, error) {

	itr := f.client.Collection("commands").OrderBy("id", firestore.Asc).Limit(1).Snapshots(context.Background())
	defer itr.Stop()
	for {
		snap, err := itr.Next()
		if err != nil {
			return nil, err
		}

		docs, err := snap.Documents.GetAll()
		if err != nil {
			return nil, err
		}
		if len(docs) == 0 {
			continue
		}

		data := docs[0].Data()
		cmd := &Cmd{
			Id:   data["id"].(int64),
			Name: data["command"].(string),
		}
		log.Printf("Firestore: Got cmd: %+v", cmd)
		return cmd, nil
	}

}

func (f *Firestore) Remove(cmd *Cmd) error {
	docs, err := f.client.Collection("commands").Where("id", "==", cmd.Id).Documents(context.Background()).GetAll()
	if err != nil {
		return err
	}
	for _, doc := range docs {
		_, err = doc.Ref.Delete(context.Background())
		if err != nil {
			return err
		}
	}
	log.Printf("Firestore: Removed cmd: id=%d", cmd.Id)
	return nil
}
