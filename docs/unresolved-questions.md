# Questions

Legend

| . | . |
| --- | --- |
| Q | question |
| C | context |
| A | answer |

1. modern error handling
    - Q: since `pkgs/errors` is archived, what's the modern equivalent of `errors.Wrap`?
    - C: 
        - learn go with tests: pointers & errors
            - read recommended article: [Don't just check for errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
                - Dave Cheney (article author) recommends using `errors.Wrap` where possible
                - `github.com/pkg/errors`'s github page says its now archived ([repo](https://github.com/pkg/errors))
    - A: 
        - in 2020, the recommendation was to use `errors.Errorf` from `github.com/pkg/errors`
        - src: https://stackoverflow.com/a/61955455/8729880
- Q: wtf is a sentinel error?
    - C:
        - was reading the following to figure out error handling best practices:
            1. asdf
            2. qwer
    - A: 
        - TODO: find answer

