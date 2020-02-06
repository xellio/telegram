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
// TODO
//
type PassportElementError interface {
}

//
// PassportElementErrorDataField - Represents an issue in one of the data fields that was provided by the user.
// The error is considered resolved when the field's value changes.
// https://core.telegram.org/bots/api#passportelementerrordatafield
// TODO
//
type PassportElementErrorDataField struct {
}

//
// PassportElementErrorFrontSide - Represents an issue with the front side of a document.
// The error is considered resolved when the file with the front side of the document changes.
// https://core.telegram.org/bots/api#passportelementerrorfrontside
// TODO
//
type PassportElementErrorFrontSide struct {
}

//
// PassportElementErrorReverseSide - Represents an issue with the reverse side of a document.
// The error is considered resolved when the file with reverse side of the document changes.
// https://core.telegram.org/bots/api#passportelementerrorreverseside
// TODO
//
type PassportElementErrorReverseSide struct {
}

//
// PassportElementErrorSelfie - Represents an issue with the selfie with a document.
// The error is considered resolved when the file with the selfie changes.
// https://core.telegram.org/bots/api#passportelementerrorselfie
// TODO
//
type PassportElementErrorSelfie struct {
}

//
// PassportElementErrorFile - Represents an issue with a document scan.
// The error is considered resolved when the file with the document scan changes.
// https://core.telegram.org/bots/api#passportelementerrorfile
// TODO
//
type PassportElementErrorFile struct {
}

//
// PassportElementErrorFiles - Represents an issue with a list of scans.
// The error is considered resolved when the list of files containing the scans changes.
// https://core.telegram.org/bots/api#passportelementerrorfiles
// TODO
//
type PassportElementErrorFiles struct {
}

//
// PassportElementErrorTranslationFile - Represents an issue with one of the files that constitute the translation of a document.
// The error is considered resolved when the file changes.
// https://core.telegram.org/bots/api#passportelementerrortranslationfile
// TODO
//
type PassportElementErrorTranslationFile struct {
}

//
// PassportElementErrorTranslationFiles - Represents an issue with the translated version of a document.
// The error is considered resolved when a file with the document translation change.
// https://core.telegram.org/bots/api#passportelementerrortranslationfiles
// TODO
//
type PassportElementErrorTranslationFiles struct {
}

//
// PassportElementErrorUnspecified - Represents an issue in an unspecified place.
// The error is considered resolved when new data is added.
// https://core.telegram.org/bots/api#passportelementerrorunspecified
// TODO
//
type PassportElementErrorUnspecified struct {
}
