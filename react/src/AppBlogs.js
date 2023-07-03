import logo from './logo.svg';
import './App.css';
import { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import { CONSTANT } from './constant';

function AppBlogs() {
  const [data, setData] = useState([]);
  useEffect(() => {
    axios.get(`${CONSTANT.URL}/api/articles`).then((data) => {
      setData(data.data.data);
    })
  }, []);
  return (
    <div className="App">
      <table className="table table-striped">
        <thead>
          <tr>
            <th scope="col">#</th>
            <th scope="col">Title</th>
            <th scope="col">Author</th>
            <th scope="col">Action</th>
          </tr>
        </thead>
        <tbody>
        {data.map((item,index) => (
          <tr key={index}>
            <th scope="row">{index+1}</th>
            <td>{item.title}</td>
            <td>{item.author}</td>
            <td><Link to={`/blog/${item.id}`} className="nav-link"><button type="button" className="btn btn-info">View</button></Link></td>
          </tr>
            ))}
        </tbody>
      </table>   

    </div>
  );
}

export default AppBlogs;
