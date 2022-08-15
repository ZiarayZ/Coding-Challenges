class SLLNode {
    constructor(value) {
        this.value = value;
        this.next;
    }

    getNode(index) {
        if (index > 1) {
            return this.next.getNode(index-1);
        }
        return this;
    }
}

class DLLNode {
    constructor(value) {
        this.value = value;
        this.prev;
        this.next;
    }
}

class SinglyLinkedList {
    constructor() {
        this.count = 0;
        this.head;
        this.tail;
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
        return false;
    }

    get popBack() {
        let secLast = this.getNode(this.count-1);
        secLast.next = NaN;
        let last = this.tail;
        this.tail = secLast;
        this.count--;
        return last.value;
    }
    get popFront() {
        let prev = this.head;
        this.head = this.head.next;
        this.count--;
        return prev.value;
    }

    getNode(index) {
        if (index > this.count) {
            return -1;//change to throw error
        }
        return this.head.getNode(index);
    }

    get values() {
        let arr = [];
        for (let i = 0; i < this.count; i++) {
            arr[i] = this.getNode(i+1).value;
        }
        return arr;
    }
}

class DoublyLinkedList {
    constructor() {
        this.count;
        this.head;
        this.tail;
    }
}

const linkedlist = new SinglyLinkedList();
linkedlist.pushBack(4);
linkedlist.pushBack(3);
linkedlist.pushBack(2);
linkedlist.pushBack(1);
console.log(linkedlist.values);