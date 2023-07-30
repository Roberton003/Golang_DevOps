package main

import  (
	"os"
	"log"
	"alertmanager"
)
func main() {
	 f, err := os.Open("alertmanager.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) }
	 if err != nil {
		 panic(err)
		 log.Fatal("error opening file: %v", err)
	 }
	 defer f.Close()
	 log.SetOutput(f)
	 log.Println("Iniciando sistema de alertas")
	 router := mux.NewRouter()
	 router.HandleFunc("/telegram", telegram.SendTelegram).Methods("POST")
	 log.Fatal(http.ListenAndServe(":8080", router))
