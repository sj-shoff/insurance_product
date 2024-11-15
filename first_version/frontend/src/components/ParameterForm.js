import React, { useState } from 'react';
import axios from 'axios';
import { TextField, Button, Box } from '@mui/material';

function ParameterForm() {
  const [name, setName] = useState('');
  const [type, setType] = useState('');
  const [defaultValue, setDefaultValue] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    const newParameter = { name, type, defaultValue };
    await axios.post('http://localhost:8000/api/parameters', newParameter);
    setName('');
    setType('');
    setDefaultValue('');
  };

  return (
    <Box component="form" onSubmit={handleSubmit} sx={{ mt: 3, mb: 3 }}>
      <TextField
        label="Name"
        value={name}
        onChange={(e) => setName(e.target.value)}
        required
        fullWidth
        margin="normal"
      />
      <TextField
        label="Type"
        value={type}
        onChange={(e) => setType(e.target.value)}
        required
        fullWidth
        margin="normal"
      />
      <TextField
        label="Default Value"
        value={defaultValue}
        onChange={(e) => setDefaultValue(e.target.value)}
        fullWidth
        margin="normal"
      />
      <Button type="submit" variant="contained" color="primary" sx={{ mt: 2 }}>
        Create Parameter
      </Button>
    </Box>
  );
}

export default ParameterForm;
