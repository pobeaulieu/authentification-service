import React, { useEffect, useState } from 'react';
import './App.css';
import { BrowserRouter, Route } from 'react-router-dom';
import Nav from './components/Nav';
import Home from './pages/Home';
import Login from './pages/Login'
import Administration from './pages/Administration';
import Business from './pages/Business';
import Residential from './pages/Residential';
import Users from './pages/Users';


function App() {
    const [loggedUser, setLoggedUser] = useState({});
    const [message, setMessage] = useState("");
    const [redirect, setRedirect] = useState(false);
    const [page, setPage] = useState(0);

    useEffect(() => {
        (
            async () => {
                const response = await fetch('http://localhost:8000/api/user', {
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
                });

                const content = await response.json();
                
                setLoggedUser(content);
                console.log(loggedUser)
            }
        )();
        // eslint-disable-next-line
    }, []);
    
    const onLogin = async (email : string, password : string) => {
        const loginResponse = await fetch('http://localhost:8000/api/login', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify({
                email,
                password
            })
        });

        const contentLogin = await loginResponse.json();

        if (contentLogin.message == "success"){
            const response = await fetch('http://localhost:8000/api/user', {
                headers: {'Content-Type': 'application/json'},
                credentials: 'include',
            });
    
            const content = await response.json();
            setLoggedUser(content);
            setRedirect(true)
        }

        else {
            setMessage(contentLogin.message)
        }
    }

    const onLogout = async () => {
        setLoggedUser({});
        await fetch('http://localhost:8000/api/logout', {
            method: 'POST',
                headers: {'Content-Type': 'application/json'},
                credentials: 'include',
            });
        setRedirect(false)
           
    }

    return (
        <div className="App">
            <BrowserRouter>
                <Nav loggedUser={loggedUser} onLogout={onLogout}/>

                <main>
                    <Route path="/" exact component={() => <Home loggedUser={loggedUser}/>}/>
                    <Route path="/login" exact component={() => <Login onLogin={onLogin} redirect={redirect} message={message}/>}/>
                    <Route path="/administration" exact component={() => <Administration loggedUser={loggedUser}/>}/>
                    <Route path="/clients/business" exact component={() => <Business loggedUser={loggedUser}/>}/>
                    <Route path="/clients/residential" exact component={() => <Residential loggedUser={loggedUser}/>}/>
                    <Route path="/users" exact component={() => <Users loggedUser={loggedUser}/>}/>
                </main>
            </BrowserRouter>
        </div>
    );
}

export default App;
