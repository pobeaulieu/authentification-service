import React, {SyntheticEvent, useState} from 'react';
import {Redirect} from "react-router-dom";

const Login = (props:any) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();
        await props.onLogin(email, password)
    }

    if (props.redirect) {
        return <Redirect to="/"/>;
    }

    return (
        <div className="form-signin">
            <form onSubmit={submit}>
                <h1 className="h3 mb-3 fw-normal">Veuillez vous connecter</h1>
                <input type="email" className="form-control" placeholder="Courriel" required
                    onChange={e => setEmail(e.target.value)}
                />

                <input type="password" className="form-control" placeholder="Mot de passe" required
                    onChange={e => setPassword(e.target.value)}
                />

                <button className="w-100 btn btn-lg btn-primary" type="submit">Se connecter</button>
                <p style={{color: 'red'}}>{props.message}</p>
            </form>
        </div>
    );
};

export default Login;