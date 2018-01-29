package bot

type GetUpdates struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	Result      []struct {
		UpdateID int      `json:"update_id"`
		Message  *Message `json:"message"`
	} `json:"result"`
}

type SendMessage struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	Result      struct {
		MessageID int    `json:"message_id"`
		From      *User  `json:"from"`
		Chat      *Chat  `json:"chat"`
		Date      int    `json:"date"`
		Text      string `json:"text"`
	} `json:"result"`
}

type KeyboardButton struct {
	Text string `json:"text"`
}

type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard"`
	OneTimeKeyboard bool               `json:"one_time_keyboard"`
}

type Message struct {
	MessageID            int              `json:"message_id"`
	From                 *User            `json:"from"` // optional
	Date                 int              `json:"date"`
	Chat                 *Chat            `json:"chat"`
	ForwardFrom          *User            `json:"forward_from"`            // optional
	ForwardFromChat      *Chat            `json:"forward_from_chat"`       // optional
	ForwardFromMessageID int              `json:"forward_from_message_id"` // optional
	ForwardDate          int              `json:"forward_date"`            // optional
	ReplyToMessage       *Message         `json:"reply_to_message"`        // optional
	EditDate             int              `json:"edit_date"`               // optional
	Text                 string           `json:"text"`                    // optional
	Entities             *[]MessageEntity `json:"entities"`                // optional
	Audio                *Audio           `json:"audio"`                   // optional
	Photo                []PhotoSize      `json:"photo"`                   // optional
	Video                *Video           `json:"video"`                   // optional
	Document             *Document        `json:"document"`                // optional

}

type User struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type GroupChat struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"`
	BigFileID   string `json:"big_file_id"`
}

type Chat struct {
	ID                  int        `json:"id"`
	Type                string     `json:"type"`
	Title               string     `json:"title"`                          // optional
	Username            string     `json:"username"`                       // optional
	FirstName           string     `json:"first_name"`                     // optional
	LastName            string     `json:"last_name"`                      // optional
	AllMembersAreAdmins bool       `json:"all_members_are_administrators"` // optional
	Photo               *ChatPhoto `json:"photo"`
	Description         string     `json:"description,omitempty"` // optional
	InviteLink          string     `json:"invite_link,omitempty"` // optional
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url"`  // optional
	User   *User  `json:"user"` // optional
}

type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"` // optional
}

type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"` // optional
	Title     string `json:"title"`     // optional
	MimeType  string `json:"mime_type"` // optional
	FileSize  int    `json:"file_size"` // optional
}

type Document struct {
	FileID    string     `json:"file_id"`
	Thumbnail *PhotoSize `json:"thumb"`     // optional
	FileName  string     `json:"file_name"` // optional
	MimeType  string     `json:"mime_type"` // optional
	FileSize  int        `json:"file_size"` // optional
}

type Sticker struct {
	FileID    string     `json:"file_id"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Thumbnail *PhotoSize `json:"thumb"`     // optional
	Emoji     string     `json:"emoji"`     // optional
	FileSize  int        `json:"file_size"` // optional
}

type Video struct {
	FileID    string     `json:"file_id"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumb"`     // optional
	MimeType  string     `json:"mime_type"` // optional
	FileSize  int        `json:"file_size"` // optional
}

type VideoNote struct {
	FileID    string     `json:"file_id"`
	Length    int        `json:"length"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumb"`     // optional
	FileSize  int        `json:"file_size"` // optional
}

type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"` // optional
	FileSize int    `json:"file_size"` // optional
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"` // optional
	UserID      int    `json:"user_id"`   // optional
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Venue struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareID string   `json:"foursquare_id"` // optional
}

type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

type File struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"` // optional
	FilePath string `json:"file_path"` // optional
}
