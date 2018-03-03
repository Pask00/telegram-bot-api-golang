package bot

import "encoding/json"

type GetUpdates struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	Result      []struct {
		UpdateID int      `json:"update_id"`
		Message  *Message `json:"message"`
	} `json:"result"`
}

type Result struct {
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message"`
}

type Response struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
}

type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	UntilDate             int64  `json:"until_date,omitempty"`
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
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
	From                 *User            `json:"from"`
	Date                 int              `json:"date"`
	Chat                 *Chat            `json:"chat"`
	ForwardFrom          *User            `json:"forward_from"`
	ForwardFromChat      *Chat            `json:"forward_from_chat"`
	ForwardFromMessageID int              `json:"forward_from_message_id"`
	ForwardDate          int              `json:"forward_date"`
	NewChatMember        *User            `json:"new_chat_member"`
	LeftChatMember       *User            `json:"left_chat_member"`
	ReplyToMessage       *Message         `json:"reply_to_message"`
	EditDate             int              `json:"edit_date"`
	Text                 string           `json:"text"`
	Entities             *[]MessageEntity `json:"entities"`
	Audio                *Audio           `json:"audio"`
	Photo                []PhotoSize      `json:"photo"`
	Video                *Video           `json:"video"`
	Document             *Document        `json:"document"`
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
	Title               string     `json:"title"`
	Username            string     `json:"username"`
	FirstName           string     `json:"first_name"`
	LastName            string     `json:"last_name"`
	AllMembersAreAdmins bool       `json:"all_members_are_administrators"`
	Photo               *ChatPhoto `json:"photo"`
	Description         string     `json:"description,omitempty"`
	InviteLink          string     `json:"invite_link,omitempty"`
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url"`
	User   *User  `json:"user"`
}

type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	MimeType  string `json:"mime_type"`
	FileSize  int    `json:"file_size"`
}

type Document struct {
	FileID    string     `json:"file_id"`
	Thumbnail *PhotoSize `json:"thumb"`
	FileName  string     `json:"file_name"`
	MimeType  string     `json:"mime_type"`
	FileSize  int        `json:"file_size"`
}

type Sticker struct {
	FileID    string     `json:"file_id"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Thumbnail *PhotoSize `json:"thumb"`
	Emoji     string     `json:"emoji"`
	FileSize  int        `json:"file_size"`
}

type Video struct {
	FileID    string     `json:"file_id"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumb"`
	MimeType  string     `json:"mime_type"`
	FileSize  int        `json:"file_size"`
}

type VideoNote struct {
	FileID    string     `json:"file_id"`
	Length    int        `json:"length"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumb"`
	FileSize  int        `json:"file_size"`
}

type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int    `json:"user_id"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Venue struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareID string   `json:"foursquare_id"`
}

type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

type File struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	FilePath string `json:"file_path"`
}
