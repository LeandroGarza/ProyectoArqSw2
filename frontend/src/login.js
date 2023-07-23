import React, { useState } from 'react';
import { Box, Button, Container, TextField, Typography } from '@mui/material';
import Cookies from "universal-cookie";
import { Link } from 'react-router-dom';
import swal from "sweetalert2";
const Cookie = new Cookies();

const Login = () => {
  const[user,setUser]= useState("");
  const[password,setPassword] = useState("");

  const onChangeUser =  (user)=>{
      setUser(user.target.value);
      
  }
  
  const onChangePas = (password)=>{
  setPassword(password.target.value)};

  
  const requestOptions={
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      
      body: JSON.stringify({username : user, password : password })
  };

  const login = async()=>{
      fetch('http://localhost:9000/login',requestOptions)
      .then(response => {if (response.status == 403) {
         swal.fire({
          text: "Datos incorrectos",
          icon: 'error',
         }).then((result) => {
          if (result.isConfirmed) {
              window.location.reload();
              return response.json()
          }})
      }
      if(response.status==200){
        swal.fire({icon: 'success'}
        ).then((result) => {
          if (result.isConfirmed) {
            window.location.replace("/")
            return response.json()
          }})
      }
      return response.json()})
      .then(response => {
          Cookie.set("user", response.token, {path: "/"})
  })
 
  };
 
  const handleSubmit= (event)=>{
      event.preventDefault();
      login();

  };

  return (
    <Container maxWidth="sm">
      <Box
        sx={{
          marginTop: 2,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <Link to="/" style={{ textDecoration: 'none' }}>
          <img
            src="https://i.pinimg.com/564x/1a/70/ed/1a70ed8c35014b66b8d87d199e5cc53b.jpg"
            alt="Logo"
            style={{ width: 195, cursor: 'pointer', marginTop: 5 }}
          />
        </Link>
        <Typography component="h2" variant="h5" >
          Registro de usuario
        </Typography>
        <Box component="form" sx={{ mt: 1 }}>
          <TextField
            label="Usuario"
            type="text"
            value={user}
            onChange={onChangeUser}
            variant="outlined"
            margin="normal"
            required
            fullWidth
          />
          <TextField
            label="Contraseña"
            type="password"
            value={password}
            onChange={onChangePas}
            variant="outlined"
            margin="normal"
            required
            fullWidth
          />
          <Button
            type="button"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2,backgroundColor: '#000000', color: '#ffffff',}}
            onClick={handleSubmit}
          >
            Iniciar Sesion
          </Button>
        </Box>
      </Box>
    </Container>
  );
};

export default Login;
