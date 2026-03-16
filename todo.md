# TODO

## Models

This are the Models that should be implemented

### MainMenu Model
- [x] Subscriptions Button
- [x] Load RSS Button

#### Load RSS

- [x] Make an input that asks for channelid and channelname and saves it into the configuration.
    - [Source](https://github.com/charmbracelet/bubbletea/blob/main/examples/textinput/main.go)

### Subscriptions Model

- [x] Browse through the Subscriptions
- [] Order By Name (asc, desc)
- [] [Pagination](https://github.com/charmbracelet/bubbletea/tree/main/examples/paginator)

### Feed Model

- [x] Show Videos
- [] Pagination


## Storage
- [x] store subscriptions into a file
- [] load that file from custom path (it can be loaded from dotfiles)  
- [] implement a singleton to access the RSS data, the videos should only be fetched once and then you fetch it from "cache".