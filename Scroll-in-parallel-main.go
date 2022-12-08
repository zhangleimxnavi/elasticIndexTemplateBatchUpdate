package main

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/olivere/elastic"
//	"golang.org/x/net/context"
//	"golang.org/x/sync/errgroup"
//	"gopkg.in/cheggaaa/pb.v2"
//	"io"
//	"time"
//)
//
////type Product struct {
////	SKU  string `json:"sku"`
////	Name string `json:"name"`
////}
//
//type SuggestField string
//type Tweet struct {
//	User     string        `json:"user"`
//	Message  string        `json:"message"`
//	Retweets int           `json:"retweets"`
//	Image    string        `json:"image,omitempty"`
//	Created  time.Time     `json:"created,omitempty"`
//	Tags     []string      `json:"tags,omitempty"`
//	Location string        `json:"location,omitempty"`
//	Suggest  *SuggestField `json:"suggest_field,omitempty"`
//}
//
//var esUrl = "http://es-cn-v0h19nq6f0009f94y.elasticsearch.aliyuncs.com:9200"
//var defaultContext = context.Background()
//
//func main() {
//	ctx := context.Background()
//	client, err := elastic.NewClient(elastic.SetURL(esUrl), elastic.SetSniff(false), elastic.SetBasicAuth("elastic", "34529@mxnavi"))
//	if err != nil {
//		panic(err)
//	}
//
//	// Count total and setup progress
//	total, err := client.Count("twitter").Type("reqstat").Do(ctx)
//	if err != nil {
//		panic(err)
//	}
//	bar := pb.StartNew(int(total))
//
//	//This example illustrates how to use goroutines to iterate
//	//through a result set via ScrollService.
//	//
//	//It uses the excellent golang.org/x/sync/errgroup package to do so.
//	//
//	//The first goroutine will Scroll through the result set and send
//	//individual documents to a channel.
//	//
//	//The second cluster of goroutines will receive documents from the channel and
//	//deserialize them.
//	//
//	//Feel free to add a third goroutine to do something with the
//	//deserialized results.
//	//
//	//Let's go.
//
//	// 1st goroutine sends individual hits to channel.
//	hits := make(chan json.RawMessage)
//	g, ctx := errgroup.WithContext(context.Background())
//	g.Go(func() error {
//		defer close(hits)
//		// Initialize scroller. Just don't call Do yet.
//		scroll := client.Scroll("twitter").Type("reqstat").Size(100)
//		for {
//			results, err := scroll.Do(ctx)
//			if err == io.EOF {
//				return nil // all results retrieved
//			}
//			if err != nil {
//				return err // something went wrong
//			}
//
//			// Send the hits to the hits channel
//			for _, hit := range results.Hits.Hits {
//				select {
//				case hits <- *hit.Source:
//				case <-ctx.Done():
//					return ctx.Err()
//				}
//			}
//		}
//		return nil
//	})
//
//	// 2nd goroutine receives hits and deserializes them.
//	//
//	// If you want, setup a number of goroutines handling deserialization in parallel.
//	for i := 0; i < 10; i++ {
//		g.Go(func() error {
//			for hit := range hits {
//				// Deserialize
//				var p Tweet
//				err := json.Unmarshal(hit, &p)
//				if err != nil {
//					return err
//				}
//
//				// Do something with the product here, e.g. send it to another channel
//				// for further processing.
//				fmt.Println(p.User, p.Message)
//				_ = p
//
//				bar.Increment()
//
//				// Terminate early?
//				select {
//				default:
//				case <-ctx.Done():
//					return ctx.Err()
//				}
//			}
//			return nil
//		})
//	}
//
//	// Check whether any goroutines failed.
//	if err := g.Wait(); err != nil {
//		panic(err)
//	}
//
//	// Done.
//	bar.Finish()
//	//bar.FinishPrint("Done")
//}
