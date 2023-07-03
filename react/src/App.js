import React, { Component, Suspense } from 'react';
import { Switch, Route, Link, Routes, BrowserRouter } from 'react-router-dom';
import AppBlogs from './AppBlogs';
import BlogDetail from './BlogDetails';


class App extends Component {
    render() {
        return (
            <BrowserRouter>
                <Suspense>
                    <Routes>
                        <Route exact path="/" element={<AppBlogs />} />
                        <Route exact path="/blog/:id" element={<BlogDetail />} />
                    </Routes>
                </Suspense>
            </BrowserRouter>
        );
    }
}

export default App;