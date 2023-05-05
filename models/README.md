# Models

Models are the core of the application. They are the conduit between source data and a vector store. 

## Anatomy of a Model

1. A model has 2 main components: a `Train` function and `Return` function.

```go
  models.Model{
    // Train is a function that returns the data to be used for training
    Train: func(ctx context.Context) []store.Vector {
      // take source data and convert it to a vector
      // then store the vector in a vector store
    },

    // Return is a function that returns the result that you want to use in your prompt
    Return: func(query string) string {
      // Return the result of the query by using the vector store
    },
},
``` 