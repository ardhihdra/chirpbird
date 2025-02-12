# ChirpBird

Our target is to build chatting-version of OmeTv 😸, so there's so much thing to be done.
Additional Docs is in /docs folder and database indices is in migration/migration.sh
so far it's just support 
- register with no login feature
- username unique check
- create group and find friends by username
- chat with them

# to be done : 
- fix front end UI/UX
- testing stuff
- continous deployment
- support server-client/end-to-end encryption
- support join and left group
- support on type and on message read
- support search and recommendation chat rooms by interest
- support random rooms
- support private rooms

# start with docker
- docker build . -t chirpbird
- docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.15.2
- sh migration.sh
- docker run -p 4000:4000 -e ES_HOST="172.17.0.1" chirpbird // where ES_HOST is the es container ip
- cd client && docker build . -t chirpbird-client
- docker run -p 3000:80 chirpbird-client

# start with docker-compose
- docker-compose up 
(notes: golang server (978879/chirpbird978879 will keep restart until the elasticsearch ready, wait ready not build yet))

# tech stack :
- front end with ReactJS
- back end with Golang and simple package
- db with Elasticsearch 😸, it's initial plan is to sync PostgreSQL and Elasticsearch with tools like
Logstash, but it will require so much time so it's only use Elastic lol (though it's not fit for every data entity), also Redis PubSub is still not used yet


Welcome to Haraj take home challenge!

In this challenge you will be assigned to help fictional startup called `ChirpBird` to create their instant messaging platform. Basically they want to create platform similar like [Facebook Messenger](https://www.messenger.com/), [WhatsApp](https://www.whatsapp.com/), or [Telegram](https://telegram.org/).

`ChirpBird` have vision to make it easy for people around the world to stay connected. It is your job to make that happen through your system. 😁

## Requirements

You need to create system with client-server architecture. For the server you need to develop it using [Go](https://golang.org/). As for the client, you just need to create a web app (you could use whatever tools or framework to create it).

The system needs to be:

1. Capable of delivering message to client instantly through Websocket
2. Capable of delivering text message
3. Capable of hosting P2P & Group chat
4. Horizontally scalable
5. Resilient from failure (since the system will be distributed, there should be tons of potential failures)

## Deployment

It is not mandatory to deploy your system on public. However if you decided to do so, it would be very great. 

At the very least you need to deploy your system using docker-compose locally.

## Evaluation

1. System design
2. Project documentation (any docs that help people understand your project better)
    - API documentation
    - Architecture diagram
    - DB schema
    - Class diagram
    - Failure handling documentation
    - etc...
3. System fail safe strategies
4. Code readibility, cleanliness, & testability
5. UI comfortability (good UX)
6. Deployment strategy (bonus)

## Submission

1. Fork this repo.
2. After you finish creating your changes, submit the link of your fork along with your CV & cover letter to [this page](https://stackoverflow.com/jobs/558729?so_medium=Talent&so_source=Talent).
3. In your cover letter, share with us what changes you have made and what further changes you would prioritize if you had more time.

## Deadline

There is no exact deadline date for this project. The only deadline is when the vacancy has been closed. We plan to open the vacancy for the next 3 months (it was first published on November 15, 2021).

So feel free to take your time.

## Questions

Got any questions? Feel free to open [issues](https://github.com/riandyrn/chirpbird/app/issues).