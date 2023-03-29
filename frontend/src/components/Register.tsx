import React, {SyntheticEvent, useState} from 'react';

const Register = (props: { onSubmit: any}) => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        await fetch('http://localhost:8000/api/register', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name,
                email,
                password
            })
        });

        props.onSubmit();
    }

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Créer un compte</h1>

            <input className="form-control" placeholder="Nom d'utilisateur" required
                   onChange={e => setName(e.target.value)}
            />

            <input type="email" className="form-control" placeholder="Courriel" required
                   onChange={e => setEmail(e.target.value)}
            />

            <input type="password" className="form-control" placeholder="Mot de passe" required
                   onChange={e => setPassword(e.target.value)}
            />

            <button className="w-100 btn btn-lg btn-primary" type="submit">Soumettre</button>
        </form>
    );
};

export default Register;