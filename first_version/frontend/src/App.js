import React from 'react';
import { Container, CssBaseline, Typography } from '@mui/material';
import ParameterForm from './components/ParameterForm';
import ParameterList from './components/ParameterList';
import RelationshipForm from './components/RelationshipForm';
import PartnerForm from './components/PartnerForm';

function App() {
  return (
    <Container component="main">
      <CssBaseline />
      <Typography component="h1" variant="h3" align="center" gutterBottom>
        Insurance Configurator
      </Typography>
      <ParameterForm />
      <ParameterList />
      <RelationshipForm />
      <PartnerForm />
    </Container>
  );
}

export default App;
