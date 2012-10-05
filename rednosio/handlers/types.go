package handlers

type Page struct {
    NavHome bool
    NavDownloads bool
    NavUploads bool
    PageTitle string
}

type BrowsePage struct {
    Page
    Files []string
}

type IndexPage struct {
    Page
    ErrMsg string
}

type RednosifyPage struct {
    Page
    ImgId string
}