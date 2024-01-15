import http from 'k6/http';
import { sleep } from 'k6';

//10 vus usuarios
//30 duration segundos de duração

export const options = {
   vus: 5000,
   duration: '30s',
};


export default function() {
   http.get('http://192.168.58.2:31500/v1/delivery/products');
   sleep(0.5);
}
