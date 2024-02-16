# Workspaces

A collection of modules on disk that are used as the main module.

When no `go.work` file exists, the workspace is made up of a single module containing the current directory.

Allows devs to work on multiple modules simultaneously.
Create a workspace by using `go work init` and add modules by using `go work use [moddir]`.


## References

- [Official Docs: Mod#workspaces](https://go.dev/ref/mod#workspaces)
- [The Go Blog: Get familiar with workspaces](https://go.dev/blog/get-familiar-with-workspaces) 
