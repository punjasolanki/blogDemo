
import './App.css';
import { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';

import { CONSTANT } from './constant'
import moment from "moment";
import { InputText } from "primereact/inputtext";
import { InputTextarea } from 'primereact/inputtextarea';



function BlogDetail() {
  const params = useParams();
  const [blogDetail, setBlogDetail] = useState([]);
  const [commentCount, setCommentCount] = useState(0);
  const [comments, setComments] = useState([]);
  const [formData, setFormData] = useState({});

  useEffect(() => {
    //GET PARTICULAR BLOG LIST
    axios.get(`${CONSTANT.URL}/api/article/${params.id}`).then((response) => {
      //console.log("BLOG DETAILS===>>>", response.data.data);
      setBlogDetail(response.data.data);
    });

    axios.get(`${CONSTANT.URL}/api/article/comments/${params.id}`).then((response) => {
      console.log("COMMENT===>>>", response.data.data);
      setCommentCount(response.data.data.length)
      setComments(response.data.data);
    });

  }, []);

  const handleValueChange = (event) => {
    const updatedValue = { ...formData };
    updatedValue[event?.target?.name] = event?.target?.value;
    setFormData(updatedValue);
  };

  const storeComment = () => {
    if (formData.name && formData.content) {
      formData.article_id = Number(params.id);
      axios.post(`${CONSTANT.URL}/api/article/comment`, formData).then((response) => {
        setComments(comment => [response.data.data, ...comment]);
        setCommentCount(commentCount + 1);
        setFormData({})
      });
    }
  }

  const removeComment = (commentId) => {
    axios.patch(`${CONSTANT.URL}/api/article/comment/${commentId}`, {}).then((response) => {
      let newArray = [...comments];
      let deletedCommentIndex = newArray.findIndex((item) => item.id == response.data.data.id)
      if (deletedCommentIndex != -1) {
        newArray.splice(deletedCommentIndex, 1);
        setComments(newArray);
        setCommentCount(commentCount - 1);
      }
    });

  }

  return (
    <>
      <div className="card-group text-center p-5">
        <div className="card">
          <div className="card-body">
            <h5 className="card-title">{blogDetail.title}</h5>
            <p className="card-text"><small className="text-muted">{blogDetail.author}</small></p>
            <p className="card-text">{blogDetail.content}</p>
          </div>
        </div>
      </div>

      <div className="card-group text-center p-5">
        <div className="card form-row">
          <div className="col-sm-12 px-2">
            <InputText
              id="name"
              name="name"
              value={formData?.name || ""}
              className="form-control mt-3"
              placeholder="Author Name"
              onChange={handleValueChange}
            />

            <InputTextarea
              id="content"
              name="content"
              placeholder="Comment..."
              className="form-control my-3"
              value={formData?.content || ""}
              onChange={handleValueChange}
              rows={5}
            //cols={50}
            />
          </div>
          <div className='m-2'>
            <button
              className={(formData.name && formData?.content) ? "btn btn-primary" : "btn btn-primary disabled"}
              onClick={storeComment}
              style={{ float: 'right' }}>Submit</button>
          </div>
        </div>
      </div>

      <div className="card-group  p-5">
        <div className="card">
          <div className="header">
            <h2>{commentCount > 0 ? "Comments" : "Comment"}   {commentCount}</h2>
          </div>
          <div className="body">
            <ul className="comment-reply list-unstyled">
              {comments.map((item, index) => (
                <li className="row clearfix p-2" key={index}>
                  <div className="icon-box col-md-1 col-4">
                    <img className="img-fluid img-thumbnail" src="https://bootdey.com/img/Content/avatar/avatar7.png" alt="Awesome Image" />
                  </div>
                  <div className="text-box col-md-10 col-8 p-l-0 p-r0">
                    <h5 className="m-b-0">{item.name}</h5>
                    <p>{item.content}</p>
                    <ul className="list-inline">
                      <li>
                        <small className="text-muted">{moment(item.timestamp).utc().format('DD-MM-YYYY hh:mm:ss')}</small>
                      </li>
                      <li>
                        <i
                          className="material-icons"
                          style={{ fontSize: '36px' }}
                          onClick={() => { removeComment(item.id) }}
                        >delete</i>
                      </li>
                    </ul>
                  </div>
                </li>
              ))}
            </ul>
          </div>
        </div>
      </div>
    </>
  );
}

export default BlogDetail;
