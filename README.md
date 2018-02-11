# XML to JSON CL Adaptor
This adaptor allows for converting of XML API's into JSON. This allows ChainLink nodes to use API's which use XML as a markup. It's built in Go using [go-json-rest](https://github.com/ant0ine/go-json-rest).

## Dependencies
To be able to run this external adaptor, you need [Docker](docker.com) installed.

## Run the Adaptor
This adaptor is published on the public docker hub, so you just need to run the two following commands:
```bash
docker pull linkpoolio/iota-cl-ea
docker run -t -p 8080:8080 linkpoolio/xml-cl-ea
```

## Using the Adaptor
To ensure that the adaptor is working properly, the following call:

> http://localhost:8080/xmltojson?endpoint=https://demo.api-platform.com/books

Should result in something similar to the following:
```json
{
"context": "/contexts/Book",
"hydra:member": [],
"hydra:totalItems": 107,
"hydra:view": {},
"id": "/books",
"type": "hydra:Collection"
}
```

## Development
To run the external adaptor, clone the repo and run the following commands:
```bash
    go get
```

To run the unit tests, run the following:
```bash
    go test
```


## Contributions
Any contributions are welcome.

Created by the [LinkPool](http://linkpool.io) Team