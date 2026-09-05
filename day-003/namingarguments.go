package main

func main() {

}

type Copier interface {
	Copy(string, string) int
}

type Copier2 interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
