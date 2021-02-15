package HttpService

import(

	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
)

type Comment = Models.Comment
// Tree struct to describe comments
type Comments struct{
	ParentComment	Comment
	ChildComments	[]Comments
}

func BFSComments(article_id int)(interface{}, error){
	comments, err := Models.GetCommentsByArticle(article_id)
	if err != nil{
		return nil, err
	}


	queue := make([]Comments, 0)
	data := make([]Comments, 0)

	for _, v := range comments{
		if v.ToCommentID == 0{
			var init_comment Comments
			init_comment.ParentComment = v 
			init_comment.ChildComments = make([]Comments, 0)
			queue = append(queue, init_comment)
		}
	}

	for _, v := range queue{
		id := v.ParentComment.ID
		var m *[]Comments
		m = &v.ChildComments
		dfs(id, comments, m)
		data = append(data, v)
	}

	return data, nil

}

func dfs(id int, comments []Comment, m *[]Comments){
	for index, value := range comments{
		if value.ToCommentID == id{
			// init Comments element
			var init_comment Comments
			init_comment.ParentComment = value 
			init_comment.ChildComments = make([]Comments, 0)
			*m = append(*m, init_comment)
			temp := *m
			new_id := value.ID
			
			// remove selected element
			comments = append(comments[:index], comments[index+1:]...)

			dfs(new_id, comments, &temp[len(temp)-1].ChildComments)
		}
	}
	return

}