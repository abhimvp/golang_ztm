# golang_ztm

- Go programming language created at google to solve google sized problems , at the end we will create a cross - platform desktop app called pixel to create and edit pixel art.
- You can find all of the course projects, code, slides in the [repository](https://github.com/jayson-lennon/ztm-golang). You can use this as a guide as you go through the course, learn Go, and build projects!
- Go is an In-demand language used by high-profile companies such as google. It's simplistic and easy to understand by design. It has built-in dependency management including a package registry similar to the javascript NPM ecosystem. it also has a familiar c style syntax.
- Technical features : It has first-class concurrency primitives and makes it perfect for backend programming such as making web servers. Has type safety which is enforced by the compiler. It's a memory safe language , accomplishes by garbage collection - no need to manage memory . Go compiles down to machine code for very fast speeds , so your programs will always run extremely quickly. Garbage collection overhead is very minimal.
- we will be working with Go CLI which is used to manage your project and we will be utilizing Go package ecosystem and it's generated documentation. Since Go is a lower level language , we'll cover concepts such as pointers, memory and concurrency.
- let's learn this language to solve a problem where we want to do some performance tuning for a company services. we're given a task to rewrite the company services to be more efficient. So this is first time optimizing applications , how can we increase the efficiency of company services? . we learn Go programming language, which was designed to build efficient services. We better get to work.

---

## Installation of Go toolchain and VSCode Extensions.

- let's go to [go.dev](https://go.dev/) & click on downloads and pick your operating system and click on that link to download the installation file.Then install.click Finish.Now your `Go toolchain` is ready to run.
- Launch VSCode -> Go to Extensions -> Install `Go` -> `After you install the Go extension, the popup should appear whenever you open a .go file. If this doesn't happen, then you can use CTRL+SHIFT+P to bring up the command window, type in go update,  check the boxes for go-outline, staticcheck, and gopls then click OK. it will then install the tools. or check all and install-all from the prompt.` In the Output bar we will see the installed things as follows:

```
2024-12-14 12:04:24.115 [info] Installing 4 tools at C:\Users\abhis\go\bin
2024-12-14 12:04:24.115 [info]   gopls
2024-12-14 12:04:24.115 [info]   goplay
2024-12-14 12:04:24.115 [info]   dlv
2024-12-14 12:04:24.115 [info]   staticcheck
2024-12-14 12:07:42.456 [info] Installing 3 tools at C:\Users\abhis\go\bin
2024-12-14 12:07:42.456 [info]   gotests
2024-12-14 12:07:42.456 [info]   gomodifytags
2024-12-14 12:07:42.456 [info]   impl
...
..
.
2024-12-14 12:08:07.992 [info] All tools successfully installed. You are ready to Go. :)
```
