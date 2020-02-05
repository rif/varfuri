package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type MapImage struct {
	Name string
	Img  string
}

var (
	maps = []*MapImage{
		{"Avram Iancu Bradatel", "avramiancu_bradatel.jpg"},
		{"Avram Iancu Dealu Livada Deal", "avramiancu_dealu_livada_deal.jpg"},
		{"Avram Iancu Dulalau", "avramiancu_dulalau.jpg"},
		{"Avram Iancu Generala", "avramiancu_generala.jpg"},
		{"Avram Iancu Glemeia", "avramiancu_glemeia.jpg"},
		{"Avram Iancu index 1", "avramiancu_index_1.jpg"},
		{"Avram Iancu index 2", "avramiancu_index_2.jpg"},
		{"Avram Iancu index 3", "avramiancu_index_3.jpg"},
		{"Avram Iancu index 4", "avramiancu_index_4.jpg"},
		{"Avram Iancu index 5", "avramiancu_index_5.jpg"},
		{"Avram Iancu index 6", "avramiancu_index_6.jpg"},
		{"Avram Iancu index 7", "avramiancu_index_7.jpg"},
		{"Avram Iancu index 8", "avramiancu_index_8.jpg"},
		{"Avram Iancu index 9", "avramiancu_index_9.jpg"},
		{"Avram Iancu Lunca", "avramiancu_lunca.jpg"},
		{"Avram Iancu Magura Tacasele", "avramiancu_magura_tacasele.jpg"},
		{"Avram Iancu Negulesti", "avramiancu_negulesti.jpg"},
		{"Avram Iancu Tacasele", "avramiancu_tacasele.jpg"},
		{"Varfurile Casoaie 1", "varfurile_casoaie_1.jpg"},
		{"Varfurile Casoaie 2", "varfurile_casoaie_2.jpg"},
		{"Varfurile Casoaie 3", "varfurile_casoaie_3.jpg"},
		{"Varfurile Glemeie 1", "varfurile_glemeie_1.jpg"},
		{"Varfurile Glemeie 2", "varfurile_glemeie_2.jpg"},
		{"Varfurile Heread", "varfurile_heread.jpg"},
		{"Varfurile Heredut", "varfurile_heredut.jpg"},
		{"Varfurile index 10", "varfurile_index_10.jpg"},
		{"Varfurile index 11", "varfurile_index_11.jpg"},
		{"Varfurile index 12", "varfurile_index_12.jpg"},
		{"Varfurile index 13", "varfurile_index_13.jpg"},
		{"Varfurile index 14", "varfurile_index_14.jpg"},
		{"Varfurile index 1", "varfurile_index_1.jpg"},
		{"Varfurile index 2", "varfurile_index_2.jpg"},
		{"Varfurile index 3", "varfurile_index_3.jpg"},
		{"Varfurile index 4", "varfurile_index_4.jpg"},
		{"Varfurile index 5", "varfurile_index_5.jpg"},
		{"Varfurile index 6", "varfurile_index_6.jpg"},
		{"Varfurile index 7", "varfurile_index_7.jpg"},
		{"Varfurile index 8", "varfurile_index_8.jpg"},
		{"Varfurile index 9", "varfurile_index_9.jpg"},
		{"Varfurile la Cris", "varfurile_la_cris.jpg"},
		{"Varfurile Magura 1", "varfurile_magura_1.jpg"},
		{"Varfurile Magura 2", "varfurile_magura_2.jpg"},
		{"Varfurile Mlaca", "varfurile_mlaca.jpg"},
		{"Varfurile Zaloage", "varfurile_zaloage.jpg"},
		{"Vidra generala", "vidra_generala.jpg"},
		{"Vidra Ilicesti", "vidra_ilicesti.jpg"},
		{"Vidra index 1", "vidra_index_1.jpg"},
		{"Vidra index 2", "vidra_index_2.jpg"},
		{"Vidra index 3", "vidra_index_3.jpg"},
		{"Vidra index 4", "vidra_index_4.jpg"},
		{"Vidra index 5", "vidra_index_5.jpg"},
		{"Vidra index 6", "vidra_index_6.jpg"},
		{"Vidra index 7", "vidra_index_7.jpg"},
		{"Vidra index 8", "vidra_index_8.jpg"},
		{"Vidra Mageresti Mihutesti", "vidra_mageresti_mihutesti.jpg"},
		{"Vidra Mageresti", "vidra_mageresti.jpg"},
		{"Vidra Mihutesti", "vidra_mihutesti.jpg"},
		{"Vidra Petrisoresti", "vidra_petrisoresti.jpg"},
		{"Vidra ceva", "vidra_ceva.jpg"},
	}
)

func contactPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("%#v", r.Form)
	if r.Method == "POST" {
		/*name := r.FormValue("from")
		email := r.FormValue("email")
		content := r.FormValue("content")
		msg := &mail.Message{
			Sender:  "Radu Fericean (Varfurile Harti) <fericean@gmail.com>",
			To:      []string{"Radu Fericean <radu@fericean.ro>"},
			Subject: fmt.Sprintf("Mesaj Harti Varfuri de la %s (%s)", name, email),
			Body:    content,
		}
		if err := mail.Send(c, msg); err != nil {
			log.Printf("Couldn't send email: %v", err)
		}*/
		return
	}
	fmt.Fprint(w, "ok")
}

func mapsPage(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(maps)
}

func main() {
	http.HandleFunc("/api/contact", contactPage)
	http.HandleFunc("/api/maps", mapsPage)
	   port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
                log.Printf("Defaulting to port %s", port)
        }

        log.Printf("Listening on port %s", port)
        if err := http.ListenAndServe(":"+port, nil); err != nil {
                log.Fatal(err)
        }

}
