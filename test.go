// // package main

// // import (
// // 	"fmt"
// // )

// // func main() {
// // 	// var int k  string j

// // 	// const text string = `ಮಾರುಕಟ್ಟೆಗೆ ಹೋಗಿ ಒಂದು ಕಿಲೋಗ್ರಾಂ ಸೇಬು ತಂದುಕೊಡಿ`
// // 	// // you can use "auto" for source language
// // 	// // so, translator will detect language
// // 	// result, _ := translategooglefree.Translate(text, "kn", "en")
// // 	// fmt.Println(result)
// // 	// // Output: "Hola, Mundo!"

// // }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"gopkg.in/oauth2.v3/errors"
// 	"gopkg.in/oauth2.v3/manage"
// 	"gopkg.in/oauth2.v3/models"
// 	"gopkg.in/oauth2.v3/server"
// 	"gopkg.in/oauth2.v3/store"
// )

// func main() {

// 	manager := manage.NewDefaultManager()
// 	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
// 	manager.MustTokenStorage(store.NewMemoryTokenStore())

// 	clientStore := store.NewClientStore()
// 	manager.MapClientStorage(clientStore)

// 	srv := server.NewDefaultServer(manager)
// 	srv.SetAllowGetAccessRequest(true)
// 	srv.SetClientInfoHandler(server.ClientFormHandler)
// 	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

// 	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
// 		log.Println("Internal Error:", err.Error())
// 		return
// 	})

// 	srv.SetResponseErrorHandler(func(re *errors.Response) {
// 		log.Println("Response Error:", re.Error.Error())
// 	})

// 	http.HandleFunc("/protected", validateToken(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello, I'm protected"))
// 	}, srv))

// 	http.HandleFunc("/credentials", func(w http.ResponseWriter, r *http.Request) {
// 		clientID := "12244555"
// 		clientSecret := "1123445667777777"

// 		err := clientStore.Set(clientID, &models.Client{
// 			ID:     clientID,
// 			Secret: clientSecret,
// 			Domain: "http://localhost:9094",
// 		})
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(map[string]string{"CLIENT_ID": clientID, "CLIENT_SECRET": clientSecret})
// 	})

// 	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
// 		srv.HandleTokenRequest(w, r)
// 	})
// 	log.Fatal(http.ListenAndServe(":9096", nil))
// }

// func validateToken(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		_, err := srv.ValidationBearerToken(r)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		f.ServeHTTP(w, r)
// 	})
// }

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {

	http.HandleFunc("/echo", handler)
	http.HandleFunc("/", home)

	fmt.Println("Test MY App")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(res http.ResponseWriter, req *http.Request) {

	Conn, err := upgrader.Upgrade(res, req, nil)

	if err != nil {
		log.Fatal("Done!!!!")
	}

	for {
		messageType, mess, err := Conn.ReadMessage()

		if err != nil {
			return
		}

		fmt.Printf("%s sent: %s\n", Conn.RemoteAddr(), string(mess))
		if err = Conn.WriteMessage(messageType, mess); err != nil {
			return
		}

	}

}

func home(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "webscoker.html")
}
