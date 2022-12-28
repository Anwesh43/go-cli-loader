## go-cli-loader

### Usage 

```
    loader := Loader{finalCapacity : 100}
    go loader.load() //todo put  go routine inside load method
    go func() {
       var BYVAL int64 = 2 
       loader.increment(BYVAL) //todo BYVAL can by int, int32 
    }()
    
```