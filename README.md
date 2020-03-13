# Golang LUIS Authoring Client
A golang client for the LUIS Authoring API. This client is generated using [goswagger](https://github.com/go-swagger/go-swagvimger). 

The source openapi specification was generated at [Programmatic API Documentation](https://australiaeast.dev.cognitive.microsoft.com/docs/services/5890b47c39e2bb17b84a55ff/operations/5890b47c39e2bb052c5b9c3d/console) and slightly modified

## Installation

Use [go get](https://golang.org/pkg/cmd/go/internal/get/) to install the client .

```bash
go get -u github.com/crazedpeanut/go-luis-authoring-client
```

## Usage

```golang
import (
	"github.com/go-openapi/strfmt"
  "github.com/go-openapi/runtime/client"
  "log"
  
  luis "github.com/crazedpeanut/go-luis-authoring-client/client"
)

func main() {
  transport := client.New("australiaeast.api.cognitive.microsoft.com", "/luis/api/v2.0", nil)
	transport.DefaultAuthentication = client.APIKeyAuth("Ocp-Apim-Subscription-Key", "header", [YOuR API KEY])
  luisAuthoringClient := luis.New(transport, strfmt.Default)
  
  params := operations.NewGetApplicationParams()
	params.SetAppID(id)
  resp, err := luisAuthoringClient.Operations.GetApplication(params, nil)

  if err != nil {
   log.Println("Your applications name is %s", resp.Payload.Name) 
  }
}

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)