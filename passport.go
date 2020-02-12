package telegram

//
// PassportData - Contains information about Telegram Passport data shared with the bot by the user.
// https://core.telegram.org/bots/api#passportdata
//
type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`
	Credentials *EncryptedCredentials       `json:"credentials"`
}

//
// PassportFile - This object represents a file uploaded to Telegram Passport.
// Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
// https://core.telegram.org/bots/api#passportfile
//
type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

//
// EncryptedPassportElement - Contains information about documents or other Telegram Passport elements shared with the bot by the user.
// https://core.telegram.org/bots/api#encryptedpassportelement
//
type EncryptedPassportElement struct {
	Type        string          `json:"type"`
	Data        string          `json:"data"`
	PhoneNumber string          `json:"phone_number"`
	Email       string          `json:"email"`
	Files       []*PassportFile `json:"files"`
	FrontSide   *PassportFile   `json:"front_side"`
	ReverseSide *PassportFile   `json:"reverse_side"`
	Selfie      *PassportFile   `json:"selfie"`
	Translation []*PassportFile `json:"translation"`
	Hash        string          `json:"hash"`
}

//
// EncryptedCredentials - Contains data required for decrypting and authenticating EncryptedPassportElement.
// See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
// https://core.telegram.org/bots/api#encryptedcredentials
//
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

//
// PassportElementError - This object represents an error in the Telegram Passport element which was submitted that should be resolved by the user.
// It should be one of:
//  - PassportElementErrorDataField
//  - PassportElementErrorFrontSide
//  - PassportElementErrorReverseSide
//  - PassportElementErrorSelfie
//  - PassportElementErrorFile
//  - PassportElementErrorFiles
//  - PassportElementErrorTranslationFile
//  - PassportElementErrorTranslationFiles
//  - PassportElementErrorUnspecified
// https://core.telegram.org/bots/api#passportelementerror
//
type PassportElementError interface {
}

//
// PassportElementErrorDataField - Represents an issue in one of the data fields that was provided by the user.
// The error is considered resolved when the field's value changes.
// https://core.telegram.org/bots/api#passportelementerrordatafield
//
type PassportElementErrorDataField struct {
	Source    string `json:"source"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	FieldName string `json:"field_name"`
	DataHash  string `json:"data_hash"`
}

//
// PassportElementErrorFrontSide - Represents an issue with the front side of a document.
// The error is considered resolved when the file with the front side of the document changes.
// https://core.telegram.org/bots/api#passportelementerrorfrontside
//
type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	Message  string `json:"message"`
	FileHash string `json:"file_hash"`
}

//
// PassportElementErrorReverseSide - Represents an issue with the reverse side of a document.
// The error is considered resolved when the file with reverse side of the document changes.
// https://core.telegram.org/bots/api#passportelementerrorreverseside
//
type PassportElementErrorReverseSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	Message  string `json:"message"`
	FileHash string `json:"file_hash"`
}

//
// PassportElementErrorSelfie - Represents an issue with the selfie with a document.
// The error is considered resolved when the file with the selfie changes.
// https://core.telegram.org/bots/api#passportelementerrorselfie
// TODO
//
type PassportElementErrorSelfie struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	Message  string `json:"message"`
	FileHash string `json:"file_hash"`
}

//
// PassportElementErrorFile - Represents an issue with a document scan.
// The error is considered resolved when the file with the document scan changes.
// https://core.telegram.org/bots/api#passportelementerrorfile
//
type PassportElementErrorFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	Message  string `json:"message"`
	FileHash string `json:"file_hash"`
}

//
// PassportElementErrorFiles - Represents an issue with a list of scans.
// The error is considered resolved when the list of files containing the scans changes.
// https://core.telegram.org/bots/api#passportelementerrorfiles
//
type PassportElementErrorFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	Message    string   `json:"message"`
	FileHashes []string `json:"file_hashes"`
}

//
// PassportElementErrorTranslationFile - Represents an issue with one of the files that constitute the translation of a document.
// The error is considered resolved when the file changes.
// https://core.telegram.org/bots/api#passportelementerrortranslationfile
//
type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	Message  string `json:"message"`
	FileHash string `json:"file_hash"`
}

//
// PassportElementErrorTranslationFiles - Represents an issue with the translated version of a document.
// The error is considered resolved when a file with the document translation change.
// https://core.telegram.org/bots/api#passportelementerrortranslationfiles
//
type PassportElementErrorTranslationFiles struct {
	Source     string `json:"source"`
	Type       string `json:"type"`
	Message    string `json:"message"`
	FileHashes string `json:"file_hashes"`
}

//
// PassportElementErrorUnspecified - Represents an issue in an unspecified place.
// The error is considered resolved when new data is added.
// https://core.telegram.org/bots/api#passportelementerrorunspecified
//
type PassportElementErrorUnspecified struct {
	Source      string `json:"source"`
	Type        string `json:"type"`
	Message     string `json:"message"`
	ElementHash string `json:"element_hash"`
}

//
// SetPassportDataErrors - Informs a user that some of the Telegram Passport elements they provided contains errors.
// The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change).
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
// https://core.telegram.org/bots/api#setpassportdataerrors
//
func (b *Bot) SetPassportDataErrors(userID int, errors []PassportElementError) (ok bool, err error) {
	params := map[string]interface{}{
		"user_id": userID,
		"errors":  errors,
	}
	result, err := b.call("setPassportDataErrors", params)
	if err != nil {
		return false, err
	}
	return result.(bool), nil
}
