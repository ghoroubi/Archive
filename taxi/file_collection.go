package main

import (
	"fmt"
	"time"
)

type File struct {
	Id            int       `json:"id"`
	File_Path     string    `json:"file_path" validate:"required"`
	Creation_Date time.Time `json:"creation_date"`
	Is_Deleted    bool      `json:"is_deleted"`
}
type FileCollection struct {
}

func (col *FileCollection) VirtualDelete(token string, id int) error {

	_, err := db.Exec(fmt.Sprintf("update file  set is_deleted=true where Id= %d", id))
	if err != nil {
		return err
	}

	return nil
}

// func (col *FileCollection) Create(f *File) (*File, error) {

// 	_, err := db.Exec(fmt.Sprintf("insert into file(file_path,creation_date) values ('%s','%s')", f.File_Path, f.Creation_Date.Format("2006-01-02 15:04:05 +0800")))

// 	if err != nil {
// 		fmt.Println("error in Create_User", err)
// 		return &File{}, err
// 	}
// 	return &File{}, nil
// }

// func (col *FileCollection) List(page int, page_Size int, token string) ([]SimpleSkill, error) {
// 	var sUsers []SimpleSkill
// 	var err error
// 	var condition, fullQuery, pagingQuery, selectQ, query string
// 	var id int
// 	id, err = usercol.Get_UserIdbytoken(token)
// 	if err != nil {
// 		fmt.Println("error in token in file", err)
// 		return nil, err
// 	}

// 	condition = " where file.is_deleted = false "
// 	if page > 0 {
// 		offset := (page - 1) * page_Size
// 		pagingQuery = " LIMIT " + strconv.Itoa(page_Size) + " OFFSET " + strconv.Itoa(offset)
// 	}
// 	selectQ = " select file.creation_date,file.file_path"
// 	query = " from file" + condition
// 	fullQuery = selectQ + query
// 	if fullQuery != "" {
// 		fullQuery += pagingQuery
// 	}

// 	err = db.Select(&sUsers, fullQuery)

// 	return sUsers, nil

// }
