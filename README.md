# ENVy
A small library used to do env fallback

## how to use
```
import github.com/hmschreck/envy
...

env_var := envy.GetEnv("KEY", "FALLBACK")
```

In the above example, if KEY exists, the value of that environment variable will be returned, else FALLBACK will be
returned