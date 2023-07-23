import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { AppBar, Toolbar, IconButton, InputBase, Button } from '@mui/material';
import { Search as SearchIcon } from '@mui/icons-material';
import '../css/home.css'
import  Cookies from 'universal-cookie';
import { red, blue, grey} from '@mui/material/colors';


const Navbar = () => {
  const cookies = new Cookies();
  const userToken = cookies.get('user');
  const navigate = useNavigate();
  const [searchTerm, setSearchTerm] = useState('');
  const handleLogout = () => {
    cookies.remove('user'); // Eliminar la cookie 'user'
    window.location.replace('/'); // Redirigir al home
  };
  

  const handleSearch = () => {
    if (searchTerm === '') {
      navigate('/');
    } else {
      const formattedSearchTerm = searchTerm.replace(' ', '%20');
      navigate(`/search/${formattedSearchTerm}`);
    }
  };

  const handleInputChange = (e) => {
    setSearchTerm(e.target.value);
  };

  return (

    <AppBar position="static" color="primary">
      <Toolbar>
      <IconButton color="inherit" onClick={handleSearch} style={{ backgroundColor: 'black' }}>
        <SearchIcon style={{ color: 'white' }} />
      </IconButton>

        <InputBase
          placeholder="Buscar"
          value={searchTerm}
          onChange={handleInputChange}
          style={{ backgroundColor: 'black', color: 'white', padding: '2px' , width: '500px' }}
        />
        {userToken ? (
          // Elementos que se mostrarán si el usuario está logueado
          <div>
            <Button onClick={handleLogout} sx={{color: '#fff', background: '#000', transform: 'scale(1)', transition: 'transform 0.3s', '&:hover': {color: red[500],transform: 'scale(1.2)'},position: 'absolute', top: '1',right: '0',justifyContent: 'flex-end', alignItems: 'flex-start', }}>
              Cerrar sesion
            </Button>
            <Link to="/create">
            <Button  sx={{color: grey[300], background: '#000', transform: 'scale(1)', transition: 'transform 0.3s', '&:hover': {color: blue[500],transform: 'scale(1.1)'}}}>
              Crear publicacion
            </Button>
            </Link>
          </div>
        ) : (
          // Elementos que se mostrarán si el usuario no está logueado
          <Link to="/login" style={{ textDecoration: 'none' }}>
           <Button style={{ backgroundColor: 'black', color: 'white' }}>
            Iniciar sesión
           </Button>
          </Link>
        )}
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;
