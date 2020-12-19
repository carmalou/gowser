# go-browse-the-web

The goal of go-browse-the-web is to build a program that can render simple HTML pages such as [motherfuckingwebsite.com](https://motherfuckingwebsite.com/). At this point, I have no plans of adding any CSS or JavaScript functionality. I am strictly focused on reading in data and then rendering the HTML content.

I decided to use Go, because ... I'm not really sure why. I guess my assumption (and it is purely an assumption) is that Go will be fast/more preformant than Node. If I have time and this isn't terrible to do, I might write it again in Node to compare.

## The current plan is as follows:

- figure out how to make a window in Go
- use curl to fetch the data
- find/use a go parser for the HTML (write my own if there's time)
- figure out the rendering (probably going to take 85% of the time)

Right now, I don't have a solid plan for dealing with the rendering. Up to now I've been 'learning' Go. I put quotes around learning because it has all been very high-level so far. I've read half of 'Introducing Go' by Caleb Doxsey, and I plan to run through the Go course on Codecademy before starting. Other than that though, it'll just be a lot of tutorials to see what I can figure out. I'll keep track of those here.
