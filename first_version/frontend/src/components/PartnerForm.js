import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { TextField, Button, Box } from '@mui/material';

function PartnerForm() {
  const [name, setName] = useState('');
  const [selectedParameters, setSelectedParameters] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    const parameterIDs = selectedParameters.split(',').map(Number);
    const newPartner = { name, parameters: parameterIDs.map(id => ({ ID: id })) };
    await axios.post('http://localhost:8000/api/partners', newPartner);
    setName('');
    setSelectedParameters('');
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
        label="Parameter IDs (comma-separated)"
        value={selectedParameters}
        onChange={(e) => setSelectedParameters(e.target.value)}
        fullWidth
        margin="normal"
      />
      <Button type="submit" variant="contained" color="primary" sx={{ mt: 2 }}>
        Create Partner
      </Button>
    </Box>
  );
}

export default PartnerForm;
