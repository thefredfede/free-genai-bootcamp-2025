import React from 'react';
import { useParams } from 'react-router-dom';

const Words: React.FC = () => {
  const { id } = useParams<{ id: string }>();

  return (
    <div>
      <h1>Words {id}</h1>
    </div>
  );
};

export default Words;