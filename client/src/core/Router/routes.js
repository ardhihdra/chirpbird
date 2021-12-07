import Login from "../../pages/login/Login.jsx";
import Dashboard from "../../pages/Dashboard.jsx";
import About from "../../pages/About";

import homeIcon from '../../assets/img/icons/023-home.png';
import newsFeed from '../../assets/img/icons/006-newsfeed.png';
import follower from '../../assets/img/icons/021-follower.png';

const menus = [
  {
    path: "/",
    exact: true,
    icon: homeIcon,
    title: 'Chirpbird',
    main: () => <Dashboard></Dashboard> 
  },
  {
    path: "/login",
    exact: true,
    icon: homeIcon,
    title: 'Chirpbird',
    main: () => <Login></Login>
  },
  {
    path: "/dashboard/:id",
    icon: newsFeed,
    title: 'Dashboard',
    main: () => <Dashboard></Dashboard> 
    //   sidebar: () => <div>shoelaces!</div>,
  },
  {
    path: "/about",
    exact: true,
    icon: follower,
    title: 'about',
    main: () => <About></About>
  },
  {
    path: "*",
    title: '404',
    main: () => {
      <main style={{ padding: "1rem" }}>
        <p>There's nothing here!</p>
      </main>
    }
  },
];

const routes = {
   menus
};

export default routes;