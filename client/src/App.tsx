import React, { useEffect, useState } from 'react';
import './App.css';
import axios from 'axios';


function App() {
  const [imageData, setImageData] = useState<string | undefined>();

  const [value, setValue] = useState<string>()

  const getDog = async () => {
    try {
      const response = await axios.post<string>(`http://localhost:8080/dogs/${value}`);
      setImageData(response.data);
      console.log(response.data);

    } catch (error) {
      console.error('Error fetching image:', error);
    }
  };

  useEffect(() => {
    getDog();
  }, [value]);

  return (
    <>
      <div style={{ marginBottom: 20, gap: 10 }}>
        <button onClick={() => setValue("affenpinscher")}>affenpinscher</button>
        <button onClick={() => setValue("dachshund")}>dachshund</button>
        <button onClick={() => setValue("pitbull")}>pitbull</button>
        <button onClick={() => setValue("labrador")}>labrador</button>
        <button onClick={() => setValue("african")}>african</button>
      </div>
      <div>
        {imageData && <img src={`data:image/jpeg;base64,${imageData}`} alt="Dog" />}
      </div>
    </>
  );
}

export default App;
