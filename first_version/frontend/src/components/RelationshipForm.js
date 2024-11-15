import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { TextField, Button, Box } from '@mui/material';

function RelationshipForm() {
  const [type, setType] = useState('');
  const [parameters, setParameters] = useState([]);
  const [selectedParameters, setSelectedParameters] = useState('');

  useEffect(() => {
    const fetchParameters = async () => {
      const result = await axios.get('http://localhost:8000/api/parameters');
      setParameters(result.data);
    };
    fetchParameters();
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const parameterIDs = selectedParameters.split(',').map(Number);
    const newRelationship = { type, parameters: parameterIDs.map(id => ({ ID: id })) };
    await axios.post('http://localhost:8000/api/relationships', newRelationship);
    setType('');
    setSelectedParameters('');
  };

  return (
    <Box component="form" onSubmit={handleSubmit} sx={{ mt: 3, mb: 3 }}>
      <TextField
        label="Type"
        value={type}
        onChange={(e) => setType(e.target.value)}
        required
        fullWidth
        margin="normal"
      />
      <TextField
        label="Parameter IDs (comma-separated)"
        value={selectedParameters}
        onChange={(e) => setSelectedParameters(e.target.value)}
        required
        fullWidth
        margin="normal"
      />
      <Button type="submit" variant="contained" color="primary" sx={{ mt: 2 }}>
        Create Relationship
      </Button>
    </Box>
  );
}

export default RelationshipForm;
