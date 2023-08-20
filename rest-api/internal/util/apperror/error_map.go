package apperror

var DBErrorMap = map[string]string{
	"email_unique":          "Email is already in use",
	"username_unique":       "Username is already in use",
	"streamkey_name_unique": "Stream key name already in use",
	"user_streamkey_unique": "Unable to generate stream key, please try again.",
}
