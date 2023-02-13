import React, { Component } from 'react'
import ReactDOM from 'react-dom'
import './index.css'

class LikeButton extends Component {
  static defaultProps = {
    likedText: '取消',
    unlikedText: '点赞'
  }
  constructor () {
    super()
    this.state = { isLiked: false }
  }
  handleClickOnButton () {
    this.setState({
      isLiked: !this.state.isLiked
    })
  }
  render () {
    return (
      <button onClick={this.handleClickOnButton.bind(this)}>
        {this.state.isLiked ? this.props.likedText : this.props.unlikedText}
      </button>
    )
  }
}

class Title extends Component {
  handleClickOnTitle (word) {
    console.log(this, word)
  }
  render () {
    return (
      <h1 onClick={this.handleClickOnTitle.bind(this, "Hello")}>React</h1>
    )
  }
}

class Header extends Component {
  renderShow (good, bad) {
    const isShow = false
    return isShow ? good : bad
  }
  render () {
    const isGoodWord = true
    return (
      <div>
        <Title />
        <h1>
          React 小书
          {
            isGoodWord
            ? <strong> is good</strong>
            : null
          }
          {
            this.renderShow(
              <strong> good word </strong>,
              <span> bad word</span>
            )
          }
        </h1>
      </div>
    )
  }
}

class Main extends Component {
  render () {
    return (
      <div>
        <h2>This is main contend</h2>
        <LikeButton likedText='赞1' unlikedText='已赞1'/>
      </div>
    )
  }
}

class Footer extends Component {
  render () {
    return (
      <div>
        <h2>This is footer</h2>
      </div>
    )
  }
}

const users = [
  { username: 'Jerry', age: 21, gender: 'male' },
  { username: 'Tomy', age: 21, gender: 'male' },
  { username: 'Lily', age: 21, gender: 'femal' },
  { username: 'Lucy', age: 21, gender: 'female' },
]

class User extends Component {
  render() {
    const { user } = this.props
    return (
      <div>
        <div>姓名：{user.username}</div>
        <div>年龄：{user.age}</div>
        <div>性别：{user.gender}</div>
        <hr />
      </div>
    )
  }
}

class Clock extends Component {
  constructor () {
    super()
    this.state = {
      data: new Date()     
    }
  }

  componentWillMount () {
    this.timer = setInterval(() => {
      this.setState({ data: new Date() })
    })
  }

  componentWillUnmount () {
    clearInterval(this.timer)
  }

  render () {
    return (
      <div>
        <h1>
          <p>现在时间是</p>
          {this.state.data.toLocaleTimeString()}
        </h1>
      </div>
    )
  }
}

class Card extends Component {
  render () {
    return (
      <div>
        <div style={{float:'left'}}>
          {this.props.children[0]}
        </div>
        <div style={{float:'right'}}>
          {this.props.children[1]}
        </div>
      </div>
    )
  }
}

class Index extends Component {
  constructor () {
    super()
    this.state = { isShowClock: true }
  }

  handleShowOrHide () {
    this.setState({
      isShowClock: !this.state.isShowClock
    })
  }

  render () {
    return (
        <div>
        <Header />
        <Main />
        <Footer />
        {users.map((user,i) => <User key={i} user={user}/>)}
        {this.state.isShowClock ? <Clock /> : null}
        <button onClick={this.handleShowOrHide.bind(this)}>
          显示或隐藏时钟
        </button>
        <Card>
          <h2>React.js 小书</h2>
          <div>开源、免费、专业、简单</div>
          订阅：<input />
        </Card>
      </div>
    )
  }
}
ReactDOM.render(
  <Index />,
  document.getElementById('root')
)