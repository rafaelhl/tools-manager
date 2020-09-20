package tools

import (
	"encoding/json"
	"strings"
	"time"

	"gorm.io/gorm"
)

type (
	Tag struct {
		ID    uint `gorm:"primary_key"`
		Value string
	}

	TagValues struct {
		values []Tag
	}

	Tool struct {
		ID          uint      `gorm:"primary_key" json:"id"`
		Title       string    `json:"title"`
		Link        string    `json:"link"`
		Description string    `json:"description"`
		Tags        TagValues `json:"tags"`

		CreatedAt time.Time      `json:"-"`
		UpdatedAt time.Time      `json:"-"`
		DeletedAt gorm.DeletedAt `json:"-"`
	}
)

func NewTags(tags ...string) TagValues {
	newTags := TagValues{}
	for _, tag := range tags {
		newTags.values = append(newTags.values, Tag{
			Value: tag,
		})
	}
	return newTags
}

func (t *TagValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Tags())
}

func (t *TagValues) String() string {
	return strings.Join(t.Tags(), ",")
}

func (t *TagValues) Tags() []string {
	tags := make([]string, len(t.values))
	for i, v := range t.values {
		tags[i] = v.Value
	}
	return tags
}

func (t *TagValues) UnmarshalJSON(data []byte) error {
	var tags []string
	err := json.Unmarshal(data, &tags)
	if err != nil {
		return err
	}

	t.values = NewTags(tags...).values
	return nil
}
