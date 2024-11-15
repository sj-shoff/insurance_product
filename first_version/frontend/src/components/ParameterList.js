import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { List, ListItem, ListItemText, Box, Typography } from '@mui/material';

function ParameterList() {
  const [parameters, setParameters] = useState([]);

  useEffect(() => {
    const fetchParameters = async () => {
      const result = await axios.get('http://localhost:8000/api/parameters');
      setParameters(result.data);
    };
    fetchParameters();
  }, []);

  return (
    <Box sx={{ mt: 3 }}>
      <Typography variant="h5" gutterBottom>
        Parameter List
      </Typography>
      <List>
        {parameters.map((parameter) => (
          <ListItem key={parameter.ID}>
            <ListItemText
              primary={parameter.Name}
              secondary={`Type: ${parameter.Type}, Default Value: ${parameter.DefaultValue}`}
            />
          </ListItem>
        ))}
      </List>
    </Box>
  );
}

export default ParameterList;
