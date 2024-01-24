package translation

import (
	"log"

	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translation "github.com/go-playground/validator/v10/translations/es"
)

func TranslationValidation(v *validator.Validate) *ut.Translator {

	es := es.New()
	uni := ut.New(es, es)

	trans, found := uni.GetTranslator("es")

	if !found {
		log.Fatal("Traductor no encontrado")
	}

	if err := es_translation.RegisterDefaultTranslations(v, trans); err != nil {
		log.Fatal(err)
	}

	return &trans
}
