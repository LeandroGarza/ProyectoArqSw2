import React, { useState, useEffect } from 'react';
import './css/home.css';
import { Card, CardContent, CardMedia, Typography, Grid } from "@mui/material";
import { Link, useParams, useNavigate } from 'react-router-dom';
import SearchIcon from '@mui/icons-material/Search';
import Navbar from './components/navbar';

const Search = () => {
  const { term } = useParams();
  const navigate = useNavigate();
  const [products, setProducts] = useState([]);

  const fetchApi = async () => {
    try {
      const url = "http://localhost:8081/search=" + term;
      const response = await fetch(url);
      const responseJSON = await response.json();
      setProducts(responseJSON);
    } catch (error) {
      console.error('Error al obtener los productos:', error);
    }
  };

  useEffect(() => {
    fetchApi();
  }, [term, navigate]);

  return (
    <div className='container'>
      <Navbar />
      <Grid container spacing={2}>
        {products && products.length > 0 ? (
          products.map((product) => (
            <Grid item xs={12} sm={6} md={4} key={product.id}>
              <Card className='product-card'>
                <CardMedia
                  component="img"
                  height="200"
                  image={product.image[0]}
                  alt={product.title[0]}
                />
                <CardContent>
                  <Typography gutterBottom variant="h5" component="div">
                    {product.title[0]}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Precio: {product.price[0]} {product.currency[0]}
                  </Typography>
                </CardContent>
                <Link to={`/product/${product.id}`} className="product-link">
                  <SearchIcon />
                  Ver detalles
                </Link>
              </Card>
            </Grid>
          ))
        ) : (
          <Typography variant="h6" component="div" align="center">
            No se encontraron resultados...
          </Typography>
        )}
      </Grid>
    </div>
  );
};

export default Search;
