package maps

import (
	//"fmt"
	"appengine"
	"appengine/mail"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var (
	maps = []string{"avramiancu_bradatel.png",
		"avramiancu_dealu_livada_deal.png",
		"avramiancu_dulalau.png",
		"avramiancu_generala.png",
		"avramiancu_glemeia.png",
		"avramiancu_index_1.png",
		"avramiancu_index_2.png",
		"avramiancu_index_3.png",
		"avramiancu_index_4.png",
		"avramiancu_index_5.png",
		"avramiancu_index_6.png",
		"avramiancu_index_7.png",
		"avramiancu_index_8.png",
		"avramiancu_index_9.png",
		"avramiancu_lunca.png",
		"avramiancu_magura_tacasele.png",
		"avramiancu_negulesti.png",
		"avramiancu_tacasele.png",
		"varfurile_casoaie_1.png",
		"varfurile_casoaie_2.png",
		"varfurile_casoaie_3.png",
		"varfurile_glemeie_1.png",
		"varfurile_glemeie_2.png",
		"varfurile_heread.png",
		"varfurile_heredut.png",
		"varfurile_index_10.png",
		"varfurile_index_11.png",
		"varfurile_index_12.png",
		"varfurile_index_13.png",
		"varfurile_index_14.png",
		"varfurile_index_1.png",
		"varfurile_index_2.png",
		"varfurile_index_3.png",
		"varfurile_index_4.png",
		"varfurile_index_5.png",
		"varfurile_index_6.png",
		"varfurile_index_7.png",
		"varfurile_index_8.png",
		"varfurile_index_9.png",
		"varfurile_la_cris.png",
		"varfurile_magura_1.png",
		"varfurile_magura_2.png",
		"varfurile_mlaca.png",
		"varfurile_zaloage.png",
		"vidra_generala.png",
		"vidra_ilicesti.png",
		"vidra_index_1.png",
		"vidra_index_2.png",
		"vidra_index_3.png",
		"vidra_index_4.png",
		"vidra_index_5.png",
		"vidra_index_6.png",
		"vidra_index_7.png",
		"vidra_index_8.png",
		"vidra_mageresti_mihutesti.png",
		"vidra_mageresti.png",
		"vidra_mihutesti.png",
		"vidra_petrisoresti.png",
		"vidra_???.png"}
)

func PrettyName(fn string)(pretty string) {
	parts := strings.Split(fn , "_")
	for _, p := range(parts){
		pretty += strings.TrimRight(p, ".png") + " "
	}
	return
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("maps").Funcs(template.FuncMap{"pretty": PrettyName}).ParseFiles("templates/base.html", "templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}	
	err = t.ExecuteTemplate(w, "base.html", map[string][]string{"Maps":maps})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func hartaPage(w http.ResponseWriter, r *http.Request) {
	harta := r.FormValue("h") 
	t, err := template.ParseFiles("templates/base.html", "templates/harta.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.Execute(w, harta)
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
			Subject: fmt.Sprintf("Mesaj Varfuri0Harti de la %s (%s)", name, email),
			Body:    content,
		}
		if err := mail.Send(c, msg); err != nil {
			c.Errorf("Couldn't send email: %v", err)
		}
		return
	}
	t, _ := template.ParseFiles("templates/base.html", "templates/contact.html")
	t.Execute(w, nil)
}

func init() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/harta", hartaPage)
	http.HandleFunc("/contact", contactPage)
}
