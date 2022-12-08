package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"gopkg.in/cheggaaa/pb.v2"
	"os"
	"sync"
)


var esUrl = "http://192.168.3.231:9200"

var defaultContext = context.Background()

var (
	r  map[string]interface{}
	wg sync.WaitGroup
)

func main() {
	client, err := elastic.NewClient(elastic.SetURL(esUrl), elastic.SetSniff(false))
	if err != nil {
		fmt.Println(err)
		return
	}

	//1、test node status

	//nodes, err := client.NodesInfo().Pretty(true).Human(true).Do(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(nodes.ClusterName)
	//fmt.Println(nodes.Nodes)

	//2、test  ping

	res, stas, err := client.Ping(esUrl).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
	fmt.Println(stas)

	//3、test index  exist?

	isExist, err := client.IndexExists("test").Do(defaultContext)
	if err != nil {
		return
	}
	fmt.Println("isExist: ", isExist)

	//4、Create a new index  for  ES 7 or later
	//ES6、 ES7 差异导致的报错， ES6 需要 type；案例上给的是ES7的案例， 无需 type
	// 参考： https://blog.csdn.net/my_miuye/article/details/110795806
	//panic: elastic: Error 400 (Bad Request): Failed to parse mapping [properties]: Root mapping definition has unsupported parameters:  [location : {type=geo_point}] [suggest_field : {type=completion}] [tags : {type=keyword}] [type=m
	//apper_parsing_exception]

	//mapping2 := `{
	//	"settings":{
	//		"number_of_shards":1,
	//		"number_of_replicas":0
	//	},
	//	"mappings":{
	//		"reqstat":{
	//			"properties":{
	//				"tags":{
	//					"type":"keyword"
	//				},
	//				"location":{
	//					"type":"geo_point"
	//				},
	//				"suggest_field":{
	//					"type":"completion"
	//				}
	//			}
	//		}
	//	}
	//}`
	//	mapping := `{
	//	"settings":{
	//		"number_of_shards":1,
	//		"number_of_replicas":0
	//	},
	//	"mappings": {
	//      "reqstat": {
	//        "dynamic": "false",
	//        "properties": {
	//          "deviceNo": {
	//            "type": "keyword"
	//          },
	//          "direction": {
	//            "type": "integer"
	//          },
	//          "domain": {
	//            "type": "keyword"
	//          },
	//          "eventTime": {
	//            "type": "long"
	//          },
	//          "productId": {
	//            "type": "keyword"
	//          },
	//          "serviceType": {
	//            "type": "keyword"
	//          },
	//          "statTs": {
	//            "type": "long"
	//          }
	//        }
	//      }
	//    }
	//}`

	//ctx := context.Background()
	////createIndex, err := client.CreateIndex("twitter").BodyString(mapping).Do(ctx)
	//createIndex, err := client.CreateIndex("twitter").Body(mapping2).Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//if !createIndex.Acknowledged {
	//	fmt.Println(createIndex.Acknowledged)
	//	// Not acknowledged
	//}
	//fmt.Println("createIndex OK")

	// 5、 delete index
	//deleteIndex, err := client.DeleteIndex("twitter").Do(defaultContext)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//if !deleteIndex.Acknowledged {
	//	fmt.Println(deleteIndex.Acknowledged)
	//}
	//fmt.Println("deleteIndex OK")

	// 6、 Add a document to an index

	//type SuggestField string
	//
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
	//// Index a tweet (using JSON serialization)
	//ctx := context.Background()
	//tweet1 := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	//put1, err := client.Index().
	//	Index("twitter").
	//	Type("reqstat").
	//	Id("1").
	//	BodyJson(tweet1).
	//	Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	//
	//// Index a second tweet (by string)
	//tweet2 := `{"user" : "olivere", "message" : "It's a Raggy Waltz"}`
	//put2, err := client.Index().
	//	Index("twitter").
	//	Type("reqstat").
	//	Id("2").
	//	BodyString(tweet2).
	//	Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//fmt.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)
	//
	//// 7、 Get document from an index
	//get1, err := client.Get().
	//	Index("twitter").
	//	Type("reqstat").
	//	Id("1").
	//	Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//var t1 Tweet
	//if get1.Found {
	//	fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	//	//  Source *json.RawMessage 是 json.RawMessage的 指针类型，获取值，需要 *get1.Source
	//	err := json.Unmarshal(*get1.Source, &t1)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println(t1.User, t1.Message, t1.Retweets, t1.Image, t1.Created, t1.Tags, t1.Location, t1.Suggest)
	//}

	// 7、Delete a document from an index
	//delRes, err := client.Delete().
	//	Index("twitter").
	//	Type("reqstat").
	//	Id("2").
	//	Do(defaultContext)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//fmt.Println(delRes.Id, delRes.Type, delRes.Index, delRes.Status, delRes.Result)
	//if delRes.Result == "deleted" {
	//	fmt.Print("Document deleted from from index\n")
	//}

	// 8、 Search for documents
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
	//// Search with a term query
	//ctx := context.Background()
	//termQuery := elastic.NewTermQuery("user", "olivere")
	////matchAllQuery := elastic.NewMatchAllQuery()
	//searchResult, err := client.Search().
	//	Index("twitter"). // search in index "twitter"
	//	Type("reqstat").
	//	Query(termQuery). // specify the query
	//	//Sort加上报错 elastic: Error 400 (Bad Request): all shards failed [type=search_phase_execution_exception] 有空研究吧
	//	//Sort("user", true).   // sort by "user" field, ascending
	//	From(0).Size(10). // take documents 0-9
	//	Pretty(true).     // pretty print request and response JSON
	//	Do(ctx)           // execute
	//if err != nil {
	//	fmt.Println(err)
	//	// Handle error
	//	//panic(err)
	//}
	//
	//// searchResult is of type SearchResult and returns hits, suggestions,
	//// and all kinds of other information from Elasticsearch.
	//fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	//
	////Each is a convenience function that iterates over hits in a search result.
	////It makes sure you don't need to check for nil values in the response.
	////However, it ignores errors in serialization. If you want full control
	////over the process, see below.
	//var ttyp Tweet
	//for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
	//	t := item.(Tweet)
	//	fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	//}
	//
	//// TotalHits is another convenience function that works even when something goes wrong.
	//fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())
	//
	//// Here's how you iterate through the search results with full control over each step.
	//fmt.Println("==============================================================")
	//if searchResult.Hits.TotalHits > 0 {
	//	fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
	//
	//	// Iterate through results
	//	for _, hit := range searchResult.Hits.Hits {
	//		// hit.Index contains the name of the index
	//
	//		// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
	//		var t Tweet
	//		err := json.Unmarshal(*hit.Source, &t)
	//		if err != nil {
	//			// Deserialization failed
	//		}
	//
	//		// Work with tweet
	//		fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	//	}
	//} else {
	//	// No hits
	//	fmt.Print("Found no tweets\n")
	//}

	//// 9、Scroll in parallel 使用 Scroll 并行处理
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
	//ctx := context.Background()
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Count total and setup progress
	//total, err := client.Count("twitter").Type("reqstat").Do(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//bar := pb.StartNew(int(total))
	//
	////This example illustrates how to use goroutines to iterate
	////through a result set via ScrollService.
	////
	////It uses the excellent golang.org/x/sync/errgroup package to do so.
	////
	////The first goroutine will Scroll through the result set and send
	////individual documents to a channel.
	////
	////The second cluster of goroutines will receive documents from the channel and
	////deserialize them.
	////
	////Feel free to add a third goroutine to do something with the
	////deserialized results.
	////
	////Let's go.
	//
	//// 1st goroutine sends individual hits to channel.
	//hits := make(chan json.RawMessage)
	//g, ctx := errgroup.WithContext(context.Background())
	//g.Go(func() error {
	//	defer close(hits)
	//	// Initialize scroller. Just don't call Do yet.
	//	scroll := client.Scroll("twitter").Type("reqstat").Size(100)
	//	for {
	//		results, err := scroll.Do(ctx)
	//		if err == io.EOF {
	//			return nil // all results retrieved
	//		}
	//		if err != nil {
	//			return err // something went wrong
	//		}
	//
	//		// Send the hits to the hits channel
	//		for _, hit := range results.Hits.Hits {
	//			select {
	//			case hits <- *hit.Source:
	//			case <-ctx.Done():
	//				return ctx.Err()
	//			}
	//		}
	//	}
	//	return nil
	//})
	//
	//// 2nd goroutine receives hits and deserializes them.
	////
	//// If you want, setup a number of goroutines handling deserialization in parallel.
	//for i := 0; i < 10; i++ {
	//	g.Go(func() error {
	//		for hit := range hits {
	//			// Deserialize
	//			var p Tweet
	//			err := json.Unmarshal(hit, &p)
	//			if err != nil {
	//				return err
	//			}
	//
	//			// Do something with the product here, e.g. send it to another channel
	//			// for further processing.
	//			fmt.Println(p.User, p.Message)
	//			_ = p
	//
	//			bar.Increment()
	//
	//			// Terminate early?
	//			select {
	//			default:
	//			case <-ctx.Done():
	//				return ctx.Err()
	//			}
	//		}
	//		return nil
	//	})
	//}
	//
	//// Check whether any goroutines failed.
	//if err := g.Wait(); err != nil {
	//	panic(err)
	//}
	//
	//// Done.
	//bar.Finish()
	////bar.FinishPrint("Done")

	//10、 Bulk indexing
	//The Bulk API in Elasticsearch enables users to perform many index/update/delete operations in a single request.
	//Instead of doing it manually (as described here)
	//you might also be interested in Bulk processor. Bulk processor does bulk processing in the background, similar to the BulkProcessor from the Java API.

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
	//tweet1 := Tweet{User: "zhanglei", Message: "go on please"}
	//tweet2 := Tweet{User: "songlaoliu", Message: "come on please"}
	//tweet3 := Tweet{User: "chengying", Message: "come on please"}
	//// Set up 4 bulk requests: 2 index requests, 1 delete request, 1 update request.
	//index1Req := elastic.NewBulkIndexRequest().Index("twitter").Type("reqstat").Id("3").Doc(tweet1)
	//index2Req := elastic.NewBulkIndexRequest().OpType("create").Index("twitter").Type("reqstat").Id("4").Doc(tweet2)
	//delete1Req := elastic.NewBulkDeleteRequest().Index("twitter").Type("reqstat").Id("1")
	//update2Req := elastic.NewBulkUpdateRequest().Index("twitter").Type("reqstat").Id("2").Doc(tweet3)
	////Doc(struct {
	////	Retweets int `json:"retweets"`
	////}{
	////	Retweets: 42,
	////})
	//
	//// Create the bulk and add the 4 requests to it
	//bulkRequest := client.Bulk()
	//bulkRequest = bulkRequest.Add(index1Req)
	//bulkRequest = bulkRequest.Add(index2Req)
	//bulkRequest = bulkRequest.Add(delete1Req)
	//bulkRequest = bulkRequest.Add(update2Req)
	//
	//// NumberOfActions contains the number of requests in a bulk
	//if bulkRequest.NumberOfActions() != 4 {
	//	fmt.Println(bulkRequest.NumberOfActions())
	//
	//}
	//
	//// Do sends the bulk requests to Elasticsearch
	//bulkResponse, err := bulkRequest.Do(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//
	//}
	//
	//// Bulk request actions get cleared
	//if bulkRequest.NumberOfActions() != 0 {
	//	fmt.Println(bulkRequest.NumberOfActions())
	//}
	//
	//// Bulk response contains valuable information about the outcome,
	//// e.g. which requests have failed etc.
	//
	//// Indexed returns information about indexed documents
	//indexed := bulkResponse.Indexed()
	//if len(indexed) != 1 {
	//	fmt.Println(len(indexed))
	//	// ...
	//}
	//if indexed[0].Id != "1" {
	//	fmt.Println(indexed[0].Id)
	//	// ...
	//}
	//if indexed[0].Status != 201 {
	//	fmt.Println(indexed[0].Status)
	//	// ...
	//}
	//
	//// Created returns information about created documents
	//created := bulkResponse.Created()
	//if len(created) != 1 {
	//	fmt.Println(len(created))
	//	// ...
	//}
	//if created[0].Id != "2" {
	//	fmt.Println(created[0].Id)
	//	// ...
	//}
	//if created[0].Status != 201 {
	//	fmt.Println(created[0].Status)
	//	// ...
	//}
	//
	//// Deleted returns information about documents that were removed
	//deleted := bulkResponse.Deleted()
	//if len(deleted) != 1 {
	//	fmt.Println(len(deleted))
	//	// ...
	//}
	//if deleted[0].Id != "1" {
	//	fmt.Println(deleted[0].Id)
	//	// ...
	//}
	//if deleted[0].Status != 200 {
	//	fmt.Println(deleted[0].Status)
	//	// ...
	//}
	////if !deleted[0].Found {
	////	// ...
	////}
	//
	//// Updated returns information about documents that were updated
	//updated := bulkResponse.Updated()
	//if len(updated) != 1 {
	//	fmt.Println(len(updated))
	//	// ...
	//}
	//if updated[0].Id != "2" {
	//	fmt.Println(updated[0].Id)
	//	// ...
	//}
	//if updated[0].Status != 200 {
	//	fmt.Println(updated[0].Status)
	//	// ...
	//}
	//if updated[0].Version != 2 {
	//	fmt.Println(updated[0].Version)
	//	// ...
	//}
	//
	//// ById returns information about documents by ID
	//id1Results := bulkResponse.ById("1")
	//if len(id1Results) != 2 {
	//	fmt.Println(len(id1Results))
	//	// Document "1" should have been indexed, then deleted
	//	// ...
	//}
	//
	//// ByAction returns information about a certain action.
	//// Use with "index", "create", "update", "delete".
	//deletedResults := bulkResponse.ByAction("delete")
	//fmt.Println(deletedResults)
	////...
	//
	//// Failed() returns information about failed bulk requests
	//// (those with a HTTP status code outside [200,299].
	//failedResults := bulkResponse.Failed()
	//fmt.Println(failedResults)

	//11、 Bulk Processor
	//A Bulk Processor is a service that can be started to receive bulk requests and commit them to Elasticsearch in the background.
	//It is similar to the BulkProcessor in the Java Client API but has some conceptual differences.

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
	//// Setup a bulk processor
	//p, err := client.BulkProcessor().
	//	Name("MyBackgroundWorker-1").
	//	Workers(2).
	//	BulkActions(1000).               // commit if # requests >= 1000
	//	BulkSize(2 << 20).               // commit if size of requests >= 2 MB
	//	FlushInterval(30 * time.Second). // commit every 30s
	//	Do(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//}
	//p.Start(context.Background())
	//
	//// Say we want to index a tweet
	//t := Tweet{User: "telexico", Message: "Elasticsearch is great."}
	//
	//// Create a bulk index request
	//// NOTICE: It is important to set the index and type here!
	//r := elastic.NewBulkIndexRequest().Index("twitter").Type("reqstat").Id("5").Doc(t)
	//
	//// Add the request r to the processor p
	//p.Add(r)
	//
	//defer p.Close()

	//12、Listing all indices
	//var skywalkingIndexs []string
	//indexNames, err := client.IndexNames()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, name := range indexNames {
	//	if strings.HasPrefix(name, "mx-sw_") {
	//		skywalkingIndexs = append(skywalkingIndexs, name)
	//	}
	//}
	//
	//for i, index := range skywalkingIndexs {
	//	fmt.Printf("index: %d,- name: %s", i, index)
	//	fmt.Println(len(skywalkingIndexs))
	//}

	//13、Listing all indices of an alias
	//aliasesResult, err := client.Aliases().Index("_all").Do(defaultContext)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//alias := aliasesResult.IndicesByAlias("index_cns3_track")
	//for i, s := range alias {
	//	fmt.Printf("index: %d,- name: %s", i, s)
	//	fmt.Println(len(s))
	//}

	// indices template exists test
	//tmpl := `{
	//  "order": 0,
	//  "index_patterns": [
	//    "mx-sw_all_p50_hour-*"
	//  ],
	//  "settings": {
	//    "index": {
	//      "refresh_interval": "120s",
	//      "number_of_shards": "2",
	//      "translog": {
	//        "sync_interval": "120s",
	//        "durability": "async"
	//      },
	//      "merge": {
	//        "scheduler": {
	//          "max_thread_count": "1"
	//        }
	//      },
	//      "max_result_window": "1000000",
	//      "analysis": {
	//        "analyzer": {
	//          "oap_analyzer": {
	//            "type": "stop"
	//          }
	//        }
	//      },
	//      "number_of_replicas": "0"
	//    }
	//  },
	//  "mappings": {
	//    "type": {
	//      "properties": {
	//        "value": {
	//          "type": "integer"
	//        },
	//        "precision": {
	//          "type": "integer"
	//        },
	//        "detail_group": {
	//          "type": "keyword"
	//        },
	//        "time_bucket": {
	//          "type": "long"
	//        }
	//      }
	//    }
	//  },
	//  "aliases": {
	//    "mx-sw_all_p50_hour": {}
	//  }
	//}`

	//putres, err := client.IndexPutTemplate("elastic-template").BodyString(tmpl).Do(context.TODO())
	//if err != nil {
	//	fmt.Printf("expected no error; got: %v", err)
	//}
	//if putres == nil {
	//	fmt.Printf("expected response; got: %v", putres)
	//}
	//if !putres.Acknowledged {
	//	fmt.Printf("expected index template to be ack'd; got: %v", putres.Acknowledged)
	//}

	// Always delete template
	//defer client.IndexDeleteTemplate("elastic-template").Do(context.TODO())

	// Check if template exists
	//exists, err := client.IndexTemplateExists("elastic-template").Do(context.TODO())
	//if err != nil {
	//	fmt.Printf("expected no error; got: %v", err)
	//}
	//if !exists {
	//	fmt.Printf("expected index template %q to exist; got: %v", "elastic-template", exists)
	//}
	//fmt.Println("the indeTemplate elastic-template exist: ", exists)

	// get all mx-sw indexs

	swTemplates, err := getAllIndexTemplatesFromFile("indexfile")
	if err != nil {
		fmt.Println(err)
		return
	}

	//for i, template := range swTemplates {
	//	fmt.Println(i, template)
	//}
	totalLength := len(swTemplates)

	//进度条长度
	bar := pb.StartNew(totalLength)

	for _, template := range swTemplates {

		//Get template body from ES
		templateResponse := getTemplateBodyFromName(template, client)

		for _, response := range templateResponse {
			//fmt.Println("key: ", k)
			//fmt.Println("value: ", response)
			updateTemplateSettings(template, response, client)
		}

		bar.Increment()
	}

	bar.Finish()

	// 单个 index template 的 测试程序
	//getres := getTemplateBodyFromName("elastic-template", client)
	//for k, response := range getres {
	//	fmt.Println("key: ", k)
	//	fmt.Println("value: ", response)
	//
	//	response.Settings["refresh_interval"] = "1000s"
	//	response.Settings["translog"] = map[string]string{
	//		"durability":    "async",
	//		"sync_interval": "2300s",
	//	}
	//
	//	fmt.Println("==================")
	//	fmt.Println(response)
	//	client.IndexPutTemplate("elastic-template").BodyJson(*response).Do(defaultContext)
	//}

}

func updateTemplateSettings(templateName string, response *elastic.IndicesGetTemplateResponse, client *elastic.Client) {
	response.Settings["refresh_interval"] = "120s"
	response.Settings["translog"] = map[string]string{
		"durability":    "async",
		"sync_interval": "240s",
	}
	//fmt.Println("==================")
	//fmt.Println(response)
	client.IndexPutTemplate(templateName).BodyJson(*response).Do(defaultContext)
}

func getTemplateBodyFromName(templateName string, client *elastic.Client) map[string]*elastic.IndicesGetTemplateResponse {
	getres, err := client.IndexGetTemplate(templateName).Do(context.TODO())
	if err != nil {
		fmt.Println(templateName)
		fmt.Printf("expected no error; got: %v", err)

	}
	if getres == nil {
		fmt.Println(templateName)
		fmt.Printf("expected to get index template %q; got: %v", templateName, getres)
	}
	return getres
}

// method for  get all mx-sw indexs
func getAllIndexTemplatesFromFile(fileName string) ([]string, error) {
	var skywalkingIndexTemplate []string

	readFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		indexName := fileScanner.Text()
		skywalkingIndexTemplate = append(skywalkingIndexTemplate, indexName)
	}

	readFile.Close()
	return skywalkingIndexTemplate, nil
}
