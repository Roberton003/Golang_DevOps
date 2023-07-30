package main

import (
    "crypto/util"
	"io/ioutil"
     "log"
    "net/http"
    "os"
	
    "github.com/gorilla/mux"
	"github.com/google/uuid"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/encrypt", Encrypt).Methods("POST")
	router.HandleFunc("/decrypt", Decrypt).Methods("POST")

	log.Printf("start http crypto server...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Decrypt(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("erro ao ler o arquivo: ", err)
		w.Header().Set("contet-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	fileName := uuid.New().String()
	tempFileName := uuid.New().String()
	
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("erro ao ler o arquivo: ", err)
		w.Header().Set("contet-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("arquivo carregado com sucesso")
	tempFile, err := ioutil.TempFile("./", fileName)
	if err != nil {
		log.Printf("erro ao criar o arquivo temporario: ", err)
		w.Header().Set("contet-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer tempFile.Close()
	log.Println("arquivo temporario criado com sucesso")

	tempFile.Write(fileBytes)
	util.DecryptLargefiles(tempFile.Name(), tempFileName, []byte("qwedrfrty12345678qwedrfrty12345678"))
	returnFileBytes, err := ioutil.ReadFile(tempFileName)
	if err != nil {
		log.Printf("erro ao ler o arquivo criptografado: ", err)
		w.Header().Set("contet-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(returnFileBytes)
	defer os.Remove(tempFileName)
	defer os.Remove(tempFile.Name())
}

func Encrypt(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("erro ao ler o arquivo: ", err)
		w.Header().Set("contet-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	fileName := uuid.New().String()
	tempFileName := uuid.New().String()

	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("erro ao ler o arquivo: ", err)
		w.Header().Set("contet-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("arquivo carregado com sucesso")
	tempFile, err := ioutil.TempFile("./", fileName)
	if err != nil {
		log.Printf("erro ao criar o arquivo temporario: ", err)
		w.Header().Set("contet-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer tempFile.Close()
	log.Println("arquivo temporario criado com sucesso")

	tempFile.Write(fileBytes)
	util.EncryptLargeFiles(tempFile.Name(), tempFileName, []byte("qwedrfrty12345678qwedrfrty12345678"))
	returnFileBytes, err := ioutil.ReadFile(tempFileName)
	log.Println("arquivo criptografado com sucesso")
	if err != nil {
		log.Printf("erro ao ler o arquivo criptografado: ", err)
		w.Header().Set("contet-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(returnFileBytes)
	defer os.Remove(tempFileName)
	defer os.Remove(tempFile.Name())
}
