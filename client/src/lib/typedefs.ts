export type MessageId = number;
export type ID = number;

export class Message {
    constructor(
        public content: string,
        public fromWho: number,
        public fromMe: boolean = false,
        public replyTo?: MessageId
    ) {}
}


export interface ApiResult {
	status: number,
	msg: string
};

export namespace Models {
	export interface Model {
		id: number;
		created_at: string;
		updated_at: string;
		deleted_at: string | undefined;
	}
	export interface MessageModel {
		id: MessageId;
		created_at: string;
		updated_at: string;
		deleted_at: string | undefined;
		content: string;
		from_id: number;
		replying_to_id: number;
	}
	export interface UserModel extends Model {
		user_name: 	string;
		nick_name: 	string | undefined;
		password:  	string;
		verified:   boolean;
		login_token: string | undefined;
	}
}
