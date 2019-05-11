package pkg

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Sub struct {
	client *firestore.Client
	datach chan map[string]interface{}
	errch  chan error
}

func New() (*Sub, error) {
	sub := &Sub{}
	sub.datach = make(chan map[string]interface{})
	sub.errch = make(chan error)

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
	sub.client = client

	startSub(sub)

	return sub, nil
}

func (s *Sub) Next() (map[string]interface{}, error) {
	select {
	case data := <-s.datach:
		return data, nil
	case err := <-s.errch:
		return nil, err
	}
}

func startSub(sub *Sub) {
	go func() {
		snapIter := sub.client.Collection("commands").Snapshots(context.Background())
		defer snapIter.Stop()

		for {
			snap, err := snapIter.Next()
			if err != nil {
				sub.errch <- err
				return
			}

			for {
				doc, err := snap.Documents.Next()
				if err != nil {
					break
				}
				sub.datach <- doc.Data()
			}
		}
	}()
}
