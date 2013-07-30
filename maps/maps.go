package maps

import (
	//"fmt"
	"appengine"
	"appengine/mail"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type MapImage struct {
	Name string
	Img  string
}

var (
	maps = []*MapImage{
		&MapImage{"Avram Iancu Bradatel", "avramiancu_bradatel.jpg"},
		&MapImage{"Avram Iancu Dealu Livada Deal", "avramiancu_dealu_livada_deal.jpg"},
		&MapImage{"Avram Iancu Dulalau", "avramiancu_dulalau.jpg"},
		&MapImage{"Avram Iancu Generala", "avramiancu_generala.jpg"},
		&MapImage{"Avram Iancu Glemeia", "avramiancu_glemeia.jpg"},
		&MapImage{"Avram Iancu index 1", "avramiancu_index_1.jpg"},
		&MapImage{"Avram Iancu index 2", "avramiancu_index_2.jpg"},
		&MapImage{"Avram Iancu index 3", "avramiancu_index_3.jpg"},
		&MapImage{"Avram Iancu index 4", "avramiancu_index_4.jpg"},
		&MapImage{"Avram Iancu index 5", "avramiancu_index_5.jpg"},
		&MapImage{"Avram Iancu index 6", "avramiancu_index_6.jpg"},
		&MapImage{"Avram Iancu index 7", "avramiancu_index_7.jpg"},
		&MapImage{"Avram Iancu index 8", "avramiancu_index_8.jpg"},
		&MapImage{"Avram Iancu index 9", "avramiancu_index_9.jpg"},
		&MapImage{"Avram Iancu Lunca", "avramiancu_lunca.jpg"},
		&MapImage{"Avram Iancu Magura Tacasele", "avramiancu_magura_tacasele.jpg"},
		&MapImage{"Avram Iancu Negulesti", "avramiancu_negulesti.jpg"},
		&MapImage{"Avram Iancu Tacasele", "avramiancu_tacasele.jpg"},
		&MapImage{"Varfurile Casoaie 1", "varfurile_casoaie_1.jpg"},
		&MapImage{"Varfurile Casoaie 2", "varfurile_casoaie_2.jpg"},
		&MapImage{"Varfurile Casoaie 3", "varfurile_casoaie_3.jpg"},
		&MapImage{"Varfurile Glemeie 1", "varfurile_glemeie_1.jpg"},
		&MapImage{"Varfurile Glemeie 2", "varfurile_glemeie_2.jpg"},
		&MapImage{"Varfurile Heread", "varfurile_heread.jpg"},
		&MapImage{"Varfurile Heredut", "varfurile_heredut.jpg"},
		&MapImage{"Varfurile index 10", "varfurile_index_10.jpg"},
		&MapImage{"Varfurile index 11", "varfurile_index_11.jpg"},
		&MapImage{"Varfurile index 12", "varfurile_index_12.jpg"},
		&MapImage{"Varfurile index 13", "varfurile_index_13.jpg"},
		&MapImage{"Varfurile index 14", "varfurile_index_14.jpg"},
		&MapImage{"Varfurile index 1", "varfurile_index_1.jpg"},
		&MapImage{"Varfurile index 2", "varfurile_index_2.jpg"},
		&MapImage{"Varfurile index 3", "varfurile_index_3.jpg"},
		&MapImage{"Varfurile index 4", "varfurile_index_4.jpg"},
		&MapImage{"Varfurile index 5", "varfurile_index_5.jpg"},
		&MapImage{"Varfurile index 6", "varfurile_index_6.jpg"},
		&MapImage{"Varfurile index 7", "varfurile_index_7.jpg"},
		&MapImage{"Varfurile index 8", "varfurile_index_8.jpg"},
		&MapImage{"Varfurile index 9", "varfurile_index_9.jpg"},
		&MapImage{"Varfurile la Cris", "varfurile_la_cris.jpg"},
		&MapImage{"Varfurile Magura 1", "varfurile_magura_1.jpg"},
		&MapImage{"Varfurile Magura 2", "varfurile_magura_2.jpg"},
		&MapImage{"Varfurile Mlaca", "varfurile_mlaca.jpg"},
		&MapImage{"Varfurile Zaloage", "varfurile_zaloage.jpg"},
		&MapImage{"Vidra generala", "vidra_generala.jpg"},
		&MapImage{"Vidra Ilicesti", "vidra_ilicesti.jpg"},
		&MapImage{"Vidra index 1", "vidra_index_1.jpg"},
		&MapImage{"Vidra index 2", "vidra_index_2.jpg"},
		&MapImage{"Vidra index 3", "vidra_index_3.jpg"},
		&MapImage{"Vidra index 4", "vidra_index_4.jpg"},
		&MapImage{"Vidra index 5", "vidra_index_5.jpg"},
		&MapImage{"Vidra index 6", "vidra_index_6.jpg"},
		&MapImage{"Vidra index 7", "vidra_index_7.jpg"},
		&MapImage{"Vidra index 8", "vidra_index_8.jpg"},
		&MapImage{"Vidra Mageresti Mihutesti", "vidra_mageresti_mihutesti.jpg"},
		&MapImage{"Vidra Mageresti", "vidra_mageresti.jpg"},
		&MapImage{"Vidra Mihutesti", "vidra_mihutesti.jpg"},
		&MapImage{"Vidra Petrisoresti", "vidra_petrisoresti.jpg"},
		&MapImage{"Vidra ceva", "vidra_ceva.jpg"},
	}
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("maps").ParseFiles("app/base.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if r.Method == "POST" {
		name := r.FormValue("from")
		email := r.FormValue("email")
		content := r.FormValue("content")
		msg := &mail.Message{
			Sender:  "Radu Fericean (Varfurile Harti) <fericean@gmail.com>",
			To:      []string{"Radu Fericean <radu@fericean.ro>"},
			Subject: fmt.Sprintf("Mesaj Harti Varfuri de la %s (%s)", name, email),
			Body:    content,
		}
		if err := mail.Send(c, msg); err != nil {
			c.Errorf("Couldn't send email: %v", err)
		}
		return
	}
	fmt.Fprint(w, "ok")
}

func mapsPage(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(maps)
}

func init() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/contact", contactPage)
	http.HandleFunc("/maps", mapsPage)
}
