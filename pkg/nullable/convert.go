package nullable

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

func NullString(nullString sql.NullString) *string {
	var value *string
	if nullString.Valid {
		value = &nullString.String
	} else {
		value = nil
	}
	return value
}

func String(stringValue *string) sql.NullString {
	if stringValue == nil {
		return sql.NullString{String: "", Valid: false}
	}

	return sql.NullString{String: *stringValue, Valid: true}
}

func NullTime(nullTime sql.NullTime) *time.Time {
	var value *time.Time
	if nullTime.Valid {
		value = &nullTime.Time
	} else {
		value = nil
	}
	return value
}

func Time(timeValue *time.Time) sql.NullTime {
	if timeValue == nil {
		return sql.NullTime{Time: time.Now(), Valid: false}
	}

	return sql.NullTime{Time: *timeValue, Valid: true}
}

func NullUUID(nullUUID uuid.NullUUID) *uuid.UUID {
	var value *uuid.UUID
	if nullUUID.Valid {
		value = &nullUUID.UUID
	} else {
		value = nil
	}
	return value
}

func UUID(uuidValue *uuid.UUID) uuid.NullUUID {
	if uuidValue == nil {
		return uuid.NullUUID{UUID: uuid.Nil, Valid: false}
	}

	return uuid.NullUUID{UUID: *uuidValue, Valid: *uuidValue != uuid.Nil}
}

func NullInt32(nullInt32 sql.NullInt32) *int32 {
	var value *int32
	if nullInt32.Valid {
		value = &nullInt32.Int32
	} else {
		value = nil
	}
	return value
}

func Int32(int32Value *int32) sql.NullInt32 {
	if int32Value == nil {
		return sql.NullInt32{Int32: 0, Valid: false}
	}

	return sql.NullInt32{Int32: *int32Value, Valid: true}
}
