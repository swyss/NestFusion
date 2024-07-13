package models

import "time"

type Task struct{
  ID string `json:"id"`
  Name string `json:"name"`
  Created time.Time `json:"created"`
  Due time.Time `json:"due"`
  Finished bool `json:"finished"`
}
