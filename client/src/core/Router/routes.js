import Home from "../../pages/Home";
import About from "../../pages/About";
import Pictures from "../../pages/Pictures";

import homeIcon from '../../assets/img/023-home.png';
import newsFeed from '../../assets/img/006-newsfeed.png';
import follower from '../../assets/img/021-follower.png';

const menus = [
  {
    path: "/",
    exact: true,
    icon: homeIcon,
    title: 'Order',
    main: () => <Home></Home>
  },
  {
    path: "/catalogue",
    icon: newsFeed,
    title: 'catalogue',
    main: () => <Pictures></Pictures> 
    //   sidebar: () => <div>shoelaces!</div>,
  },
  {
    path: "/about",
    exact: true,
    icon: follower,
    title: 'about',
    main: () => <About></About>
  },
];

const routes = {
   menus
};

export default routes;