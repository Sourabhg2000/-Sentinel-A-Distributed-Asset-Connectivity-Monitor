import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import { createPinia } from 'pinia';

// PrimeVue Core and Theme
import PrimeVue from 'primevue/config';
import Aura from '@primeuix/themes/aura';

// PrimeVue Components
import Button from 'primevue/button';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Toast from 'primevue/toast';
import ConfirmDialog from 'primevue/confirmdialog'; // Add this for your sessions page

// 1. Import the Services
import ConfirmationService from 'primevue/confirmationservice';
import ToastService from 'primevue/toastservice';

const app = createApp(App);

// 2. Use Services (Crucial step to fix your error)
app.use(ConfirmationService);
app.use(ToastService);

app.use(router);
app.use(createPinia());

app.use(PrimeVue, {
    theme: {
        preset: Aura
    }
});

// 3. Register Global Components
app.component('Button', Button);
app.component('DataTable', DataTable);
app.component('Toast', Toast);
app.component('Column', Column);
app.component('ConfirmDialog', ConfirmDialog); // Register globally for convenience

app.mount('#app');