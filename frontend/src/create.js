import React, { useState, useEffect } from 'react';
import { Box, Button, Container, TextField, Typography } from '@mui/material';
import Cookies from "universal-cookie";
import swal from "sweetalert2";
const Cookie = new Cookies();

const Create = () => {
  const [token, setToken] = useState("");
  const [title, setTitle] = useState("");
  const [price, setPrice] = useState(0);
  const [currency, setCurrency] = useState("");
  const [image, setImage] = useState("");
  const [state, setState] = useState("");
  const [address, setAddress] = useState("")
  const [condition, setCondition] = useState("")

  useEffect(() => {
    loadTokenFromCookie();
  }, []);

  const loadTokenFromCookie = () => {
    const userToken = Cookie.get("user");
    setToken(userToken);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    // Verificar si algún campo está vacío
    if (
      !title ||
      !price ||
      !currency ||
      !image ||
      !state ||
      !address ||
      !condition
    ) {
      swal.fire({
        text: "Debes rellenar todos los campos",
        icon: 'error'
      });
      return; // Detener la ejecución si hay campos vacíos
    }

    const requestOptions = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify([
        {
          title,
          price,
          currency,
          image,
          state,
          address,
          condition,
        }
      ])
    };

    try {
      const response = await fetch('http://localhost:8090/items', requestOptions);
      if (response.status === 201) {
        swal.fire({
          text: "Publicación creada exitosamente",
          icon: 'success'
        }).then(() => {
          setTitle("");
          setPrice(0);
          setCurrency("");
          setImage("");
          setState("");
          setAddress(0);
          setCondition("");
        });
      } else {
        swal.fire({
          text: "Error al crear la publicación",
          icon: 'error'
        });
      }
    } catch (error) {
      swal.fire({
        text: "Error en la solicitud",
        icon: 'error'
      });
    }
  };

  return (
    <Container maxWidth="sm">
      <Box
        sx={{
          marginTop: 4,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <Typography component="h2" variant="h5">
          Postular un inmueble
        </Typography>
        <Box component="form" sx={{ mt: 1 }}>
          <TextField label="Título" type="text" value={title} onChange={(event) => setTitle(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Precio" type="number" value={price} onChange={(event) => setPrice(Number(event.target.value))} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Moneda" type="text" value={currency} onChange={(event) => setCurrency(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Imagen (URL)" type="url" value={image} onChange={(event) => setImage(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Estado" type="number" value={state} onChange={(event) => setState(Number(event.target.value))} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Direccion" type="text" value={address} onChange={(event) => setAddress(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Condicion" type="text" value={condition} onChange={(event) => setCondition(event.target.value)} variant="outlined" margin="normal" required fullWidth />

          <Button type="button" fullWidth variant="contained" sx={{color: '#fff', background: '#000'}} onClick={handleSubmit}>
            Crear Publicación
          </Button>
        </Box>
      </Box>
    </Container>
  );
};

export default Create;
