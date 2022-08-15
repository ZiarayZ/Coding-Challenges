class SLLNode {
    constructor(value) {
        this.value = value;
        this.next;
    }

    getNextNode(index) {
        if (index > 0) {
            return this.next.getNextNode(index-1);
        }
        return this;
    }
}

class DLLNode extends SLLNode {
    constructor(value) {
        super(value);
        this.prev;
    }

    getPrevNode(index) {
        if (index > 0) {
            return this.prev.getPrevNode(index-1);
        }
        return this;
    }
}

class LinkedList {
    constructor() {
        this.count = 0;
        this.head;
        this.tail;
    }

    get popFront() {
        let previous = this.head;
        this.head = this.head.next;
        this.count--;
        return previous.value;
    }

    contains(item) {
        for (let i = 0; i < this.count; i++) {
            if (this.getNode(i).value === item) {
                return true;//could also use a integer location instead
            }
        }
        return false;//which would mean returning -1 here
    }

    get values() {
        let arr = [];
        for (let i = 0; i < this.count; i++) {
            arr[i] = this.getNode(i).value;
        }
        return arr;
    }
}

class SinglyLinkedList extends LinkedList {
    constructor() {
        super();
    }

    pushBack(value) {
        this.count++;
        if (this.count === 1) {
            this.head = new SLLNode(value);
            this.tail = this.head;
            return true;
        } else if (this.count === 2) {
            this.tail = new SLLNode(value);
            this.head.next = this.tail;
            return true;
        } else {
            this.tail.next = new SLLNode(value);
            this.tail = this.tail.next;
            return true;
        }
        this.count--;
        return false;
    }
    pushFront(value) {
        this.count++;
        if (this.count === 1) {
            this.head = new SLLNode(value);
            this.tail = this.head;
            return true;
        } else if (this.count === 2) {
            this.head = new SLLNode(value);
            this.head.next = this.tail;
            return true;
        } else {
            let second = this.head;
            this.head = new SLLNode(value);
            this.head.next = second;
            return true;
        }
        this.count--;
        return false;
    }

    get popBack() {
        let secLast = this.getNode(this.count-2);
        secLast.next = NaN;
        let last = this.tail;
        this.tail = secLast;
        this.count--;
        return last.value;
    }

    remove(index) {
        if (index >= this.count || index < 0) {
            return false;//change to throw index error
        } else if (index === 0) {
            this.head = this.head.next;
            this.count--;
            return true;
        } else if (index > 0) {
            let item = this.getNode(index-1);
            item.next = item.next.next;
            this.count--;
            return true;
        }

        return false;//change to throw type error
    }

    getNode(index) {
        if (index >= this.count || index < 0) {
            return -1;//change to throw error
        }
        return this.head.getNextNode(index);
    }
}

class DoublyLinkedList extends LinkedList {
    constructor() {
        super();
    }

    pushBack(value) {
        this.count++;
        if (this.count === 1) {
            this.head = new DLLNode(value);
            this.tail = this.head;
            return true;
        } else if (this.count === 2) {
            this.tail = new DLLNode(value);
            this.tail.prev = this.head;
            this.head.next = this.tail;
            return true;
        } else {
            this.tail.next = new DLLNode(value);
            this.tail.next.prev = this.tail;
            this.tail = this.tail.next;
            return true;
        }
        this.count--;
        return false;
    }
    pushFront(value) {
        this.count++;
        if (this.count === 1) {
            this.head = new DLLNode(value);
            this.tail = this.head;
            return true;
        } else if (this.count === 2) {
            this.head = new DLLNode(value);
            this.head.next = this.tail;
            this.tail.prev = this.head;
            return true;
        } else {
            this.head.prev = new DLLNode(value);
            this.head.prev.next = this.head;
            this.head = this.head.prev;
            return true;
        }
        this.count--;
        return false;
    }
    
    get popBack() {
        let following = this.tail;
        this.tail = this.tail.prev;
        this.count--;
        return following.value;
    }

    remove(index) {
        if (index >= this.count || index < 0) {
            return false;//change to throw index error
        } else if (index === 0) {
            this.head = this.head.next;
            this.count--;
            return true;
        } else if (index === this.count-1) {
            this.tail = this.tail.prev;
            this.count--;
            return true;
        } else if (index > 0) {
            let item = this.getNode(index);
            item.next.prev = item.prev;
            item.prev.next = item.next;
            item = item.next;
            this.count--;
            return true;
        }

        return false;//change to throw type error
    }

    getNode(index) {
        if (index >= this.count) {
            return -1;//change to throw index error
        } else if (index < 0) {
            if (index <= -this.count) {
                return -1;//change to throw index error
            } else if (index >= -this.count/2) {
                return this.tail.getPrevNode(-index);
            } else if (index <= -this.count/2) {
                return this.head.getNextNode(1-this.count + index);
            }
        } else if (index <= this.count/2) {
            return this.head.getNextNode(index);
        } else if (index >= this.count/2) {
            return this.tail.getPrevNode(this.count-1 - index);
        }

        return -1;//throw type error
    }
}

const linkedlist = new SinglyLinkedList();
linkedlist.pushBack(4);
linkedlist.pushBack(3);
linkedlist.pushBack(2);
linkedlist.pushBack(1);
console.log(linkedlist.values);

console.log("");

linkedlist = new DoublyLinkedList();
linkedlist.pushBack(4);
linkedlist.pushBack(3);
linkedlist.pushBack(2);
linkedlist.pushBack(1);
console.log(linkedlist.values);