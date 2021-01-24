package HttpService

import(

	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
)

type Comment = Models.Comment
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
	for _, v := range comments{
		if v.ToCommentID == id{
			var init_comment Comments
			init_comment.ParentComment = v 
			init_comment.ChildComments = make([]Comments, 0)
			*m = append(*m, init_comment)
			temp := *m
			new_id := v.ID
			dfs(new_id, comments, &temp[len(temp)-1].ChildComments)
		}
	}
	return

}