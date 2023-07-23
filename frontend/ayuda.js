const [searchInput, setSearchInput] = useState('');
const [filteredProducts, setFilteredProducts] = useState(products);

const handleSearchInputChange = (event) => {
  setSearchInput(event.target.value);
};

const handleSearchButtonClick = () => {
  const searchKeywords = searchInput.split(" ").filter(keyword => keyword !== ""); // Obtener las palabras clave separadas
  const filteredProducts = products.filter((product) =>
    searchKeywords.every(keyword =>
      product.description.toLowerCase().includes(keyword.toLowerCase())
    )
  );
  setFilteredProducts(filteredProducts);
};

const handleLogoClick = () => {
  setSearchInput('');
  setFilteredProducts(products);
};

<div className='container'>
<div className='barra'>
   <Link to="/" onClick={handleLogoClick} className='rompebolas'>
     <img src='https://cdn.freebiesupply.com/logos/large/2x/jl-logo-black-and-white.png' className='loguito'/>
   </Link>     
     <Link to="/login">
       <button className="login-button">Login</button>
     </Link>
   <div className='busqueda'>
     <input
       type="text"
       placeholder="Buscar..."
       className="search-input"
       onChange={handleSearchInputChange}
     />
     <button className='lupa' onClick={handleSearchButtonClick}>
       <SearchIcon />
     </button>
   </div>
 </div>
     <Grid container spacing={2}>
       {filteredProducts.map((product) => (
         <Grid item key={product.id} xs={12} sm={6} md={4}>
           <Link to={`/${product.id}`} className="product-link">
             <Card sx={{margin:1, height: "89%", backgroundColor: '#5c77cf', transition: 'background-color 0.3s', '&:hover': {backgroundColor: '#8094d8'}}} className="product-card" >
               <CardMedia 
                 component="img"
                 height="200"
                 image={product.picture}
                 alt={product.id}
               />
               <CardContent className='product-content'>
                 <Typography variant="h6" component="div">
                   {product.title}
                 </Typography>
                 <Typography variant="subtitle1" color="textSecondary">
                   {product.seller}
                 </Typography>
                 <Typography variant="body2" color="textSecondary">
                   Precio: {product.price} {product.currency}
                 </Typography>
                 <Typography variant="body2" component="p">
                   Descripción: {truncateDescription(product.description, 15)}
                 </Typography>
                 <Typography variant="body2" color="textSecondary">
                   Estado: {product.state}
                 </Typography>
                 <Typography variant="body2" color="textSecondary">
                   Ciudad: {product.city}
                 </Typography>
                 <Typography variant="body2" color="textSecondary">
                   Calle: {product.street} {product.number}
                 </Typography>
               </CardContent>
             </Card>
           </Link>
         </Grid>
       ))}
     </Grid>
   </div>